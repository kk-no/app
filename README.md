# experiment application

This repository is experimental.  
Don't use too many.

## Dependencies

- GCP
  - Workload identity pools
  - Cloud Run
- Go

## Commands

```
$ gcloud iam workload-identity-pools create "my-pool" \
  --project="${PROJECT}" \
  --location="global" \
  --display-name="Test pool"

$ gcloud iam workload-identity-pools providers create-oidc "my-provider" \
  --project="${PROJECT}" \
  --location="global" \
  --workload-identity-pool="my-pool" \
  --display-name="Test provider" \
  --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.aud=assertion.aud" \
  --issuer-uri="https://token.actions.githubusercontent.com"``

$ gcloud iam service-accounts create sa

$ gcloud iam service-accounts add-iam-policy-binding "sa@${PROJECT}.iam.gserviceaccount.com" \
  --project="${PROJECT}" \
  --role="roles/iam.workloadIdentityUser" \
  --member="principalSet://iam.googleapis.com/projects/${PROJECT_ID}/locations/global/workloadIdentityPools/github/attribute.repository/${GH_USER}/${REPO}"
```