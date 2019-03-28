gorun:
	GOOGLE_APPLICATION_CREDENTIALS="${HOME}/Documents/gcloud/webapp/accounts/webapp-github_svc-account-dev-1_c70e2fbaaa2e.json" go run . 
godev:
	dev_appserver.py .
load:
	hey -n 2000 -c 100 -m GET -H "Authorization: Bearer okok" -t 2 "http://localhost:8080/v1/accounts/123"
