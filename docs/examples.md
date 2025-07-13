# Examples

Here are some real-world examples of using CDEvents CLI to automate CI/CD workflows.

## Jenkins Pipeline

Integrate CDEvents CLI in a Jenkins pipeline to track build progress:

```groovy
pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                // Generate a pipeline started event with custom data
                sh '''docker run --rm cdevents-cli:latest send --target http://your-events-endpoint.com pipeline started \
                    --id "${BUILD_ID}" --name "${JOB_NAME}" \
                    --custom "build_number=${BUILD_NUMBER}" \
                    --custom "branch=${BRANCH_NAME}" \
                    --custom "commit=${GIT_COMMIT}" \
                    --custom "executor=${EXECUTOR_NUMBER}"'''

                // Execute your build steps
                sh 'make build'

                // Generate a build finished event with build metrics
                sh '''docker run --rm cdevents-cli:latest send --target http://your-events-endpoint.com build finished \
                    --id "${BUILD_ID}" --name "${JOB_NAME}" --outcome "success" \
                    --custom-json '{
                        "data": {
                            "duration": "${BUILD_DURATION}",
                            "artifacts": ["app.jar", "tests.xml"],
                            "test_results": {
                                "passed": 150,
                                "failed": 0,
                                "skipped": 5
                            }
                        },
                        "labels": {
                            "team": "backend",
                            "environment": "ci"
                        }
                    }' '''
            }
        }
    }

    post {
        success {
            // Generate a pipeline finished event on success
            sh '''docker run --rm cdevents-cli:latest send --target http://your-events-endpoint.com pipeline finished \
                --id "${BUILD_ID}" --name "${JOB_NAME}" --outcome "success" \
                --custom "total_duration=${BUILD_DURATION}" \
                --custom "workspace=${WORKSPACE}"'''
        }

        failure {
            // Generate a pipeline finished event on failure with error details
            sh '''docker run --rm cdevents-cli:latest send --target http://your-events-endpoint.com pipeline finished \
                --id "${BUILD_ID}" --name "${JOB_NAME}" --outcome "failure" \
                --errors "${BUILD_LOG}" \
                --custom "failed_stage=${env.STAGE_NAME}" \
                --custom "error_type=build_failure"'''
        }
    }
}
```

## GitHub Actions

Use CDEvents CLI within GitHub Actions to emit events:

```yaml
name: CI Pipeline

on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3

    - name: Generate Pipeline Started Event
      run: |
        docker run --rm cdevents-cli:latest send --target http://events.example.com pipeline started --id "$GITHUB_RUN_ID" --name "$GITHUB_WORKFLOW"

    - name: Run Build
      run: make build

    - name: Generate Build Finished Event
      run: |
        docker run --rm cdevents-cli:latest send --target http://events.example.com build finished --id "$GITHUB_RUN_ID" --name "$GITHUB_WORKFLOW" --outcome success

    - name: Generate Pipeline Finished Event
      if: always()
      run: |
        outcome="failure"
        if [ "${{ job.status }}" == "success" ]; then outcome="success"; fi
        docker run --rm cdevents-cli:latest send --target http://events.example.com pipeline finished --id "$GITHUB_RUN_ID" --name "$GITHUB_WORKFLOW" --outcome "$outcome"
```

## GitLab CI

Integrate CDEvents CLI in GitLab CI to report the status of deployments:

```yaml
stages:
  - build
  - deploy

build_job:
  stage: build
  script:
    - docker run --rm cdevents-cli:latest send --target http://events.example.com pipeline started --id "$CI_PIPELINE_ID" --name "$CI_PROJECT_NAME"
    - make build
    - docker run --rm cdevents-cli:latest send --target http://events.example.com build finished --id "$CI_PIPELINE_ID" --name "$CI_PROJECT_NAME" --outcome success

deploy_job:
  stage: deploy
  script:
    - docker run --rm cdevents-cli:latest send --target http://events.example.com service deployed --id "$CI_PIPELINE_ID" --name "$CI_PROJECT_NAME" --environment production
```

## Real-Time Monitoring with HTTP

Leverage HTTP endpoints to monitor event flow:

```bash
# Start a local HTTP server
python3 -m http.server 8080

# Send events to the server
cdevents-cli send --target http://localhost:8080 pipeline started --id "pipeline-123" --name "my-pipeline"
```

## Testing Events Locally

Test without sending real events:

```bash
# Output event to console in CloudEvent format
cdevents-cli generate pipeline started --id "pipeline-123" --name "my-pipeline" --output cloudevent

# Save event to a file
cdevents-cli generate pipeline started --id "pipeline-123" --name "my-pipeline" --output json > event.json
```

## Advanced Custom Data Examples

### Kubernetes Deployment Events

```bash
# Service deployment with Kubernetes metadata
cdevents-cli generate service deployed --id "my-app-v1.2.3" --name "my-app" \
  --custom-yaml '
data:
  version: "1.2.3"
  namespace: "production"
  replicas: 3
  image: "my-app:1.2.3"
  resources:
    requests:
      cpu: "100m"
      memory: "128Mi"
    limits:
      cpu: "500m"
      memory: "512Mi"
labels:
  app: "my-app"
  version: "v1.2.3"
  environment: "production"
  team: "platform"
annotations:
  deployment.kubernetes.io/revision: "42"
  prometheus.io/scrape: "true"
  prometheus.io/port: "8080"
links:
  - name: "deployment"
    url: "https://k8s.example.com/api/v1/namespaces/production/deployments/my-app"
    type: "kubernetes"
  - name: "pods"
    url: "https://k8s.example.com/api/v1/namespaces/production/pods?labelSelector=app=my-app"
    type: "kubernetes"
  - name: "monitoring"
    url: "https://grafana.example.com/d/app-dashboard/my-app?var-namespace=production"
    type: "monitoring"
'
```

### Test Results with Detailed Metrics

```bash
# Test suite finished with comprehensive results
cdevents-cli generate test testsuite-finished --id "test-suite-123" --name "Integration Tests" \
  --outcome "success" \
  --custom-json '{
    "data": {
      "duration_seconds": 145,
      "total_tests": 250,
      "passed": 245,
      "failed": 0,
      "skipped": 5,
      "flaky": 2,
      "coverage": {
        "line": 87.5,
        "branch": 82.3,
        "function": 95.1
      },
      "performance": {
        "avg_response_time": 45,
        "p95_response_time": 89,
        "p99_response_time": 156
      }
    },
    "labels": {
      "test_type": "integration",
      "environment": "staging",
      "browser": "chrome",
      "platform": "linux"
    },
    "annotations": {
      "test.framework": "pytest",
      "test.parallel": "true",
      "test.retry.enabled": "true"
    },
    "links": [
      {
        "name": "test-report",
        "url": "https://ci.example.com/job/123/test-report",
        "type": "report"
      },
      {
        "name": "coverage-report",
        "url": "https://ci.example.com/job/123/coverage",
        "type": "coverage"
      },
      {
        "name": "artifacts",
        "url": "https://ci.example.com/job/123/artifacts",
        "type": "artifacts"
      }
    ]
  }'
```

### Build Event with Artifact Information

```bash
# Build finished with artifact details
cdevents-cli generate build finished --id "build-456" --name "Frontend Build" \
  --outcome "success" \
  --custom-json '{
    "data": {
      "duration_seconds": 180,
      "artifacts": [
        {
          "name": "app.js",
          "size_bytes": 1024000,
          "checksum": "sha256:abc123...",
          "url": "https://cdn.example.com/builds/456/app.js"
        },
        {
          "name": "styles.css",
          "size_bytes": 51200,
          "checksum": "sha256:def456...",
          "url": "https://cdn.example.com/builds/456/styles.css"
        }
      ],
      "dependencies": {
        "npm_packages": 247,
        "vulnerabilities": {
          "high": 0,
          "medium": 2,
          "low": 5
        }
      },
      "build_tools": {
        "webpack": "5.88.0",
        "babel": "7.22.0",
        "typescript": "5.1.0"
      }
    },
    "labels": {
      "build_type": "production",
      "target": "web",
      "optimization": "true"
    },
    "links": [
      {
        "name": "build-logs",
        "url": "https://ci.example.com/build/456/logs",
        "type": "logs"
      },
      {
        "name": "security-scan",
        "url": "https://security.example.com/scan/456",
        "type": "security"
      }
    ]
  }'
```
