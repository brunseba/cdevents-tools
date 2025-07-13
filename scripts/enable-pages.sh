#!/bin/bash

# Script to enable GitHub Pages for the repository
# This script uses the GitHub CLI to enable Pages with GitHub Actions

echo "🔧 Enabling GitHub Pages for the repository..."

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
    echo "❌ GitHub CLI (gh) is not installed. Please install it first:"
    echo "   https://github.com/cli/cli#installation"
    exit 1
fi

# Check if authenticated
if ! gh auth status &> /dev/null; then
    echo "❌ Not authenticated with GitHub CLI. Please run: gh auth login"
    exit 1
fi

# Get repository information
REPO=$(gh repo view --json name,owner --jq '.owner.login + "/" + .name')
echo "📂 Repository: $REPO"

# Enable Pages with GitHub Actions
echo "🚀 Enabling GitHub Pages with GitHub Actions source..."

# Use GitHub API to enable Pages
gh api \
  --method POST \
  -H "Accept: application/vnd.github+json" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  /repos/$REPO/pages \
  -f source='{"branch":"main","path":"/"}' \
  -f build_type='workflow' \
  2>/dev/null || echo "ℹ️  Pages may already be enabled"

# Check Pages status
echo "📊 Checking Pages status..."
PAGES_STATUS=$(gh api /repos/$REPO/pages --jq '.status // "Not found"' 2>/dev/null)
PAGES_URL=$(gh api /repos/$REPO/pages --jq '.html_url // "Not available"' 2>/dev/null)

if [ "$PAGES_STATUS" = "Not found" ]; then
    echo "❌ Failed to enable Pages. Please enable manually:"
    echo "   1. Go to repository Settings → Pages"
    echo "   2. Set Source to 'GitHub Actions'"
    echo "   3. Save the settings"
else
    echo "✅ GitHub Pages enabled successfully!"
    echo "📝 Status: $PAGES_STATUS"
    echo "🌐 URL: $PAGES_URL"
    echo ""
    echo "🎉 Your documentation will be available at: $PAGES_URL"
    echo "⏳ It may take a few minutes for the site to become available after the first deployment."
fi

echo ""
echo "📋 Next steps:"
echo "1. The workflow will automatically deploy documentation on pushes to main"
echo "2. Check the Actions tab for deployment status"
echo "3. Visit the Pages URL once deployment is complete"
