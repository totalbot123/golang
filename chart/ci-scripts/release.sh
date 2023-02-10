#!/bin/bash

set -ex

##
# Step 1: Generate new version
##
previous_version=$(semversioner current-version)
semversioner release
new_version=$(semversioner current-version)

##
# Step 2: Generate CHANGELOG.md
##
echo "Generating CHANGELOG.md file..."
semversioner changelog > CHANGELOG.md
# Use new version in the README.md examples
echo "Replace version '$previous_version' to '$new_version' in README.md ..."
sed -i "s/$previous_version/$new_version/g" $BITBUCKET_REPO_SLUG/Chart.yaml
# Use new version in the pipe.yml metadata file

##
# Step 3: Generate README.md
##
docker run --rm --volume "$(pwd)/deployment-chart:/helm-docs" -u $(id -u) jnorwood/helm-docs:latest
mv $(pwd)/deployment-chart/README.md $BITBUCKET_CLONE_DIR

##
# Step 4: Build and push docker image
##
echo $ARTIFACT_WRITER_KEY > /opt/gcp_key.json
gcloud auth activate-service-account --key-file=/opt/gcp_key.json
helm package $BITBUCKET_REPO_SLUG
helm push $BITBUCKET_REPO_SLUG-*.tgz oci://$ARTIFACT_REGION-docker.pkg.dev/$ARTIFACT_PROJECT/$BITBUCKET_REPO_SLUG

##
# Step 5: Commit back to the repository
##
echo "Committing updated files to the repository..."
git add --all
git reset -- $BITBUCKET_REPO_SLUG-*.tgz
git commit -m "Update files for new version '${new_version}' [skip ci]"
git push origin ${BITBUCKET_BRANCH}

##
# Step 6: Tag the repository
##
echo "Tagging for release ${new_version}"
git tag -a -m "Tagging for release ${new_version}" "${new_version}"
git push origin ${new_version}
