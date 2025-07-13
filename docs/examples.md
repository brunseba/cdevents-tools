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
                    --custom-json '{
                        "data": {
                            "build_number": "${BUILD_NUMBER}",
                            "branch": "${BRANCH_NAME}",
                            "commit": "${GIT_COMMIT}",
                            "executor": "${EXECUTOR_NUMBER}"
                        },
                        "labels": {
                            "team": "backend",
                            "environment": "ci"
                        }
                    }' '''

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
                --custom-json '{
                    "data": {
                        "total_duration": "${BUILD_DURATION}",
                        "workspace": "${WORKSPACE}"
                    },
                    "labels": {
                        "result": "success"
                    }
                }' '''
        }

        failure {
            // Generate a pipeline finished event on failure with error details
            sh '''docker run --rm cdevents-cli:latest send --target http://your-events-endpoint.com pipeline finished \
                --id "${BUILD_ID}" --name "${JOB_NAME}" --outcome "failure" \
                --errors "${BUILD_LOG}" \
                --custom-json '{
                    "data": {
                        "failed_stage": "${env.STAGE_NAME}",
                        "error_type": "build_failure"
                    },
                    "labels": {
                        "result": "failure"
                    }
                }' '''
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
  --custom-json '{
    "data": {
      "version": "1.2.3",
      "namespace": "production",
      "replicas": 3,
      "image": "my-app:1.2.3",
      "resources": {
        "requests": {
          "cpu": "100m",
          "memory": "128Mi"
        },
        "limits": {
          "cpu": "500m",
          "memory": "512Mi"
        }
      }
    },
    "labels": {
      "app": "my-app",
      "version": "v1.2.3",
      "environment": "production",
      "team": "platform"
    },
    "annotations": {
      "deployment.kubernetes.io/revision": "42",
      "prometheus.io/scrape": "true",
      "prometheus.io/port": "8080"
    },
    "links": [
      {
        "name": "deployment",
        "url": "https://k8s.example.com/api/v1/namespaces/production/deployments/my-app",
        "type": "kubernetes"
      },
      {
        "name": "pods",
        "url": "https://k8s.example.com/api/v1/namespaces/production/pods?labelSelector=app=my-app",
        "type": "kubernetes"
      },
      {
        "name": "monitoring",
        "url": "https://grafana.example.com/d/app-dashboard/my-app?var-namespace=production",
        "type": "monitoring"
      }
    ]
  }'
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

## CLI Usage Examples with Input and Output

Here are detailed examples showing the exact CLI commands and their outputs in different formats.

### Basic Pipeline Event Generation

#### Input:
```bash
cdevents-cli generate pipeline started --id "pipeline-123" --name "my-build-pipeline" --source "jenkins-ci" --url "https://ci.example.com/job/my-build-pipeline/123"
```

#### Output (JSON - default):
```json
{
  "context": {
    "version": "0.4.1",
    "id": "01234567-89ab-cdef-0123-456789abcdef",
    "source": "jenkins-ci",
    "type": "dev.cdevents.pipelinerun.started.0.2.0",
    "timestamp": "2024-01-15T10:30:00Z"
  },
  "subject": {
    "id": "pipeline-123",
    "source": "jenkins-ci",
    "type": "pipelineRun",
    "content": {
      "pipelineName": "my-build-pipeline",
      "url": "https://ci.example.com/job/my-build-pipeline/123"
    }
  }
}
```

### Pipeline Event with Custom Data

#### Input:
```bash
cdevents-cli generate pipeline finished --id "pipeline-456" --name "frontend-deploy" \
  --outcome "success" \
  --source "github-actions" \
  --custom-json '{
    "data": {
      "duration_seconds": 120,
      "commit_sha": "abc123def456",
      "branch": "main",
      "environment": "production",
      "deployment_strategy": "blue-green"
    },
    "labels": {
      "team": "frontend",
      "project": "e-commerce",
      "criticality": "high"
    },
    "annotations": {
      "ci.tool": "github-actions",
      "deploy.tool": "kubernetes",
      "monitoring.enabled": "true"
    },
    "links": [
      {
        "name": "deployment-logs",
        "url": "https://github.com/myorg/frontend/actions/runs/123456",
        "type": "logs"
      },
      {
        "name": "live-site",
        "url": "https://app.example.com",
        "type": "application"
      }
    ]
  }'
```

#### Output (JSON):
```json
{
  "context": {
    "version": "0.4.1",
    "id": "87654321-dcba-9876-5432-109876543210",
    "source": "github-actions",
    "type": "dev.cdevents.pipelinerun.finished.0.2.0",
    "timestamp": "2024-01-15T10:32:00Z"
  },
  "subject": {
    "id": "pipeline-456",
    "source": "github-actions",
    "type": "pipelineRun",
    "content": {
      "pipelineName": "frontend-deploy",
      "outcome": "success",
      "url": ""
    }
  },
  "customData": {
    "data": {
      "duration_seconds": 120,
      "commit_sha": "abc123def456",
      "branch": "main",
      "environment": "production",
      "deployment_strategy": "blue-green"
    },
    "labels": {
      "team": "frontend",
      "project": "e-commerce",
      "criticality": "high"
    },
    "annotations": {
      "ci.tool": "github-actions",
      "deploy.tool": "kubernetes",
      "monitoring.enabled": "true"
    },
    "links": [
      {
        "name": "deployment-logs",
        "url": "https://github.com/myorg/frontend/actions/runs/123456",
        "type": "logs"
      },
      {
        "name": "live-site",
        "url": "https://app.example.com",
        "type": "application"
      }
    ]
  },
  "customDataContentType": "application/json"
}
```

### Build Event with Different Output Formats

#### Input:
```bash
cdevents-cli generate build finished --id "build-789" --name "api-service" \
  --outcome "failure" \
  --errors "Unit tests failed: 3 failures out of 150 tests" \
  --source "gitlab-ci" \
  --url "https://gitlab.example.com/backend/api-service/-/jobs/789" \
  --custom-json '{
    "data": {
      "test_results": {
        "total": 150,
        "passed": 147,
        "failed": 3,
        "skipped": 0
      },
      "coverage_percent": 82.5,
      "build_duration_seconds": 45
    },
    "labels": {
      "team": "backend",
      "service": "api",
      "language": "go"
    }
  }' \
  --output yaml
```

#### Output (YAML):
```yaml
context:
  version: "0.4.1"
  id: "fedcba98-7654-3210-fedc-ba9876543210"
  source: "gitlab-ci"
  type: "dev.cdevents.build.finished.0.2.0"
  timestamp: "2024-01-15T10:35:00Z"
subject:
  id: "build-789"
  source: "gitlab-ci"
  type: "build"
  content:
    errors: "Unit tests failed: 3 failures out of 150 tests"
    outcome: "failure"
    url: "https://gitlab.example.com/backend/api-service/-/jobs/789"
customData:
  data:
    test_results:
      total: 150
      passed: 147
      failed: 3
      skipped: 0
    coverage_percent: 82.5
    build_duration_seconds: 45
  labels:
    team: "backend"
    service: "api"
    language: "go"
customDataContentType: "application/json"
```

### Test Event with CloudEvent Output

#### Input:
```bash
cdevents-cli generate test testcase-finished --id "test-001" --name "integration-tests" \
  --outcome "success" \
  --source "pytest-runner" \
  --url "https://ci.example.com/test-reports/001" \
  --custom-json '{
    "data": {
      "test_suite": "integration",
      "duration_ms": 15000,
      "assertions": 45,
      "database_calls": 12,
      "api_calls": 8
    },
    "labels": {
      "environment": "staging",
      "database": "postgresql",
      "browser": "chrome"
    }
  }' \
  --output cloudevent
```

#### Output (CloudEvent):
```json
{
  "specversion": "1.0",
  "id": "abcdef12-3456-7890-abcd-ef1234567890",
  "source": "pytest-runner",
  "type": "dev.cdevents.testcaserun.finished.0.2.0",
  "subject": "test-001",
  "time": "2024-01-15T10:38:00Z",
  "datacontenttype": "application/json",
  "data": {
    "context": {
      "version": "0.4.1",
      "id": "abcdef12-3456-7890-abcd-ef1234567890",
      "source": "pytest-runner",
      "type": "dev.cdevents.testcaserun.finished.0.2.0",
      "timestamp": "2024-01-15T10:38:00Z"
    },
    "subject": {
      "id": "test-001",
      "source": "pytest-runner",
      "type": "testCaseRun",
      "content": {
        "environment": {
          "id": "staging",
          "source": "pytest-runner"
        },
        "outcome": "success",
        "testCase": {
          "id": "integration-tests",
          "name": "integration-tests",
          "type": "integration",
          "uri": "https://ci.example.com/test-reports/001"
        }
      }
    },
    "customData": {
      "data": {
        "test_suite": "integration",
        "duration_ms": 15000,
        "assertions": 45,
        "database_calls": 12,
        "api_calls": 8
      },
      "labels": {
        "environment": "staging",
        "database": "postgresql",
        "browser": "chrome"
      }
    },
    "customDataContentType": "application/json"
  }
}
```

### Service Deployment Event

#### Input:
```bash
cdevents-cli generate service deployed --id "service-prod-v2.1.0" --name "payment-service" \
  --environment "production" \
  --source "argocd" \
  --url "https://argocd.example.com/applications/payment-service" \
  --custom-json '{
    "data": {
      "version": "v2.1.0",
      "image": "registry.example.com/payment-service:v2.1.0",
      "replicas": 5,
      "namespace": "production",
      "resources": {
        "cpu_request": "500m",
        "memory_request": "1Gi",
        "cpu_limit": "1000m",
        "memory_limit": "2Gi"
      },
      "health_check": {
        "readiness_probe": "/health/ready",
        "liveness_probe": "/health/live",
        "startup_probe": "/health/startup"
      }
    },
    "labels": {
      "team": "payments",
      "criticality": "critical",
      "compliance": "pci-dss",
      "monitoring": "prometheus"
    },
    "annotations": {
      "deployment.strategy": "rolling-update",
      "security.scan": "passed",
      "performance.test": "passed"
    },
    "links": [
      {
        "name": "kubernetes-deployment",
        "url": "https://k8s.example.com/api/v1/namespaces/production/deployments/payment-service",
        "type": "kubernetes"
      },
      {
        "name": "monitoring-dashboard",
        "url": "https://grafana.example.com/d/payment-service/payment-service-dashboard",
        "type": "monitoring"
      },
      {
        "name": "api-documentation",
        "url": "https://docs.example.com/api/payment-service/v2.1.0",
        "type": "documentation"
      }
    ]
  }'
```

#### Output (JSON):
```json
{
  "context": {
    "version": "0.4.1",
    "id": "12345678-90ab-cdef-1234-567890abcdef",
    "source": "argocd",
    "type": "dev.cdevents.service.deployed.0.2.0",
    "timestamp": "2024-01-15T10:40:00Z"
  },
  "subject": {
    "id": "service-prod-v2.1.0",
    "source": "argocd",
    "type": "service",
    "content": {
      "artifactId": "",
      "environment": {
        "id": "production",
        "source": "argocd",
        "uri": "https://argocd.example.com/applications/payment-service"
      }
    }
  },
  "customData": {
    "data": {
      "version": "v2.1.0",
      "image": "registry.example.com/payment-service:v2.1.0",
      "replicas": 5,
      "namespace": "production",
      "resources": {
        "cpu_request": "500m",
        "memory_request": "1Gi",
        "cpu_limit": "1000m",
        "memory_limit": "2Gi"
      },
      "health_check": {
        "readiness_probe": "/health/ready",
        "liveness_probe": "/health/live",
        "startup_probe": "/health/startup"
      }
    },
    "labels": {
      "team": "payments",
      "criticality": "critical",
      "compliance": "pci-dss",
      "monitoring": "prometheus"
    },
    "annotations": {
      "deployment.strategy": "rolling-update",
      "security.scan": "passed",
      "performance.test": "passed"
    },
    "links": [
      {
        "name": "kubernetes-deployment",
        "url": "https://k8s.example.com/api/v1/namespaces/production/deployments/payment-service",
        "type": "kubernetes"
      },
      {
        "name": "monitoring-dashboard",
        "url": "https://grafana.example.com/d/payment-service/payment-service-dashboard",
        "type": "monitoring"
      },
      {
        "name": "api-documentation",
        "url": "https://docs.example.com/api/payment-service/v2.1.0",
        "type": "documentation"
      }
    ]
  },
  "customDataContentType": "application/json"
}
```

### Task Event with Minimal Data

#### Input:
```bash
cdevents-cli generate task started --id "task-security-scan" --name "security-vulnerability-scan" \
  --pipeline "pipeline-security-checks" \
  --source "sonarqube" \
  --url "https://sonar.example.com/project/overview?id=my-project"
```

#### Output (JSON):
```json
{
  "context": {
    "version": "0.4.1",
    "id": "98765432-10ab-cdef-9876-543210abcdef",
    "source": "sonarqube",
    "type": "dev.cdevents.taskrun.started.0.2.0",
    "timestamp": "2024-01-15T10:42:00Z"
  },
  "subject": {
    "id": "task-security-scan",
    "source": "sonarqube",
    "type": "taskRun",
    "content": {
      "taskName": "security-vulnerability-scan",
      "pipelineRun": {
        "id": "pipeline-security-checks",
        "source": "sonarqube"
      },
      "url": "https://sonar.example.com/project/overview?id=my-project"
    }
  }
}
```

### Sending Events to Different Targets

#### Input: Send to HTTP endpoint
```bash
cdevents-cli send --target "https://events.example.com/webhook" \
  --headers "Authorization=Bearer your-token-here" \
  --headers "Content-Type=application/json" \
  --retries 3 \
  --timeout 30s \
  pipeline finished --id "pipeline-deploy-prod" --name "production-deployment" \
  --outcome "success" \
  --source "jenkins" \
  --custom-json '{
    "data": {
      "deployment_id": "deploy-12345",
      "version": "v1.5.2",
      "duration_seconds": 180
    },
    "labels": {
      "environment": "production",
      "team": "devops"
    }
  }'
```

#### Output:
```
Event sent to https://events.example.com/webhook: pipeline-deploy-prod
```

#### Input: Send to file
```bash
cdevents-cli send --target "file://events.json" \
  build started --id "build-123" --name "microservice-build" \
  --source "github-actions" \
  --url "https://github.com/myorg/microservice/actions/runs/123"
```

#### Output:
```
Event sent to file events.json: build-123
```

#### Contents of events.json:
```json
{
  "context": {
    "version": "0.4.1",
    "id": "11111111-2222-3333-4444-555555555555",
    "source": "github-actions",
    "type": "dev.cdevents.build.started.0.2.0",
    "timestamp": "2024-01-15T10:45:00Z"
  },
  "subject": {
    "id": "build-123",
    "source": "github-actions",
    "type": "build",
    "content": {
      "url": "https://github.com/myorg/microservice/actions/runs/123"
    }
  }
}
```
