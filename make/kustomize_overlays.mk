
##@ Kustomize Overlay Generation

## Targets to help create deployment kustomizations (overlays)

CLUSTER_NAME ?= $(KIND_CLUSTER_NAME)

DEPLOYMENT_COUNT ?= 2
DEPLOYMENT_NAMESPACE ?= dns-operator
DEPLOYMENT_NAME_SUFFIX ?= 1
DEPLOYMENT_WATCH_NAMESPACES ?=

GCP_CREDENTIALS_FILE ?= config/local-setup/dns-provider/gcp/gcp-credentials.env
AWS_CREDENTIALS_FILE ?= config/local-setup/dns-provider/aws/aws-credentials.env
AZURE_CREDENTIALS_FILE ?= config/local-setup/dns-provider/azure/azure-credentials.env

## Location to generate cluster overlays
CLUSTER_OVERLAY_DIR ?= $(shell pwd)/tmp/overlays
$(CLUSTER_OVERLAY_DIR):
	mkdir -p $(CLUSTER_OVERLAY_DIR)

.PHONY: generate-cluster-overlay
generate-cluster-overlay: remove-cluster-overlay ## Generate a cluster overlay with namespaced deployments for the current cluster (CLUSTER_NAME)
	# Generate cluster overlay
	mkdir -p $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)
	cd $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME) && \
	touch kustomization.yaml && \
	$(KUSTOMIZE) edit add resource "../../../config/crd"

	# Generate common dns provider kustomization
	mkdir -p $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/dns-providers
	cd $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/dns-providers && \
	touch kustomization.yaml && \
	$(KUSTOMIZE) edit add secret dns-provider-credentials-inmemory --disableNameSuffixHash --from-literal=INMEM_INIT_ZONES=kuadrant.local --type "kuadrant.io/inmemory"

	# Add dns providers that require credentials
	@if [[ -f $(GCP_CREDENTIALS_FILE) ]]; then\
		cp config/local-setup/dns-provider/gcp/gcp-credentials.env $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/dns-providers/ ;\
		cd $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/dns-providers && \
		$(KUSTOMIZE) edit add secret dns-provider-credentials-gcp --disableNameSuffixHash --from-env-file=gcp-credentials.env --type "kuadrant.io/gcp" ;\
	fi
	@if [[ -f $(AWS_CREDENTIALS_FILE) ]]; then\
		cp config/local-setup/dns-provider/aws/aws-credentials.env $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/dns-providers/ ;\
		cd $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/dns-providers && \
		$(KUSTOMIZE) edit add secret dns-provider-credentials-aws --disableNameSuffixHash --from-env-file=aws-credentials.env --type "kuadrant.io/aws" ;\
	fi
	@if [[ -f $(AZURE_CREDENTIALS_FILE) ]]; then\
		cp config/local-setup/dns-provider/azure/azure-credentials.env $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/dns-providers/ ;\
		cd $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/dns-providers && \
		$(KUSTOMIZE) edit add secret dns-provider-credentials-azure --disableNameSuffixHash --from-env-file=azure-credentials.env --type "kuadrant.io/azure" ;\
	fi

	@n=1 ; while [[ $$n -le $(DEPLOYMENT_COUNT) ]] ; do \
		$(MAKE) -s generate-operator-deployment-overlay DEPLOYMENT_NAME_SUFFIX=$$n DEPLOYMENT_NAMESPACE=${DEPLOYMENT_NAMESPACE}-$$n DEPLOYMENT_WATCH_NAMESPACES=${DEPLOYMENT_NAMESPACE}-$$n ;\
		cd $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME) && $(KUSTOMIZE) edit add resource namespace-${DEPLOYMENT_NAMESPACE}-$$n && cd - > /dev/null ;\
		((n = n + 1)) ;\
	done ;\

.PHONY: remove-cluster-overlay
remove-cluster-overlay: ## Remove an existing cluster overlay for the current cluster (CLUSTER_NAME)
	rm -rf $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)

.PHONY: remove-all-cluster-overlays
remove-all-cluster-overlays: ## Remove all existing cluster overlays (kuadrant-dns-local*)
	rm -rf $(CLUSTER_OVERLAY_DIR)/kuadrant-dns-local*

.PHONY: generate-operator-deployment-overlay
generate-operator-deployment-overlay: ## Generate a DNS Operator deployment overlay for the current cluster (CLUSTER_NAME)
	# Generate dns-operator deployment overlay
	mkdir -p $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/namespace-$(DEPLOYMENT_NAMESPACE)/dns-operator
	cd $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/namespace-$(DEPLOYMENT_NAMESPACE)/dns-operator && \
	touch kustomization.yaml && \
	$(KUSTOMIZE) edit add resource "../../../../../config/local-setup/dns-operator" && \
	$(KUSTOMIZE) edit set namesuffix -- -$(DEPLOYMENT_NAME_SUFFIX)  && \
	$(KUSTOMIZE) edit add patch --kind Deployment --patch '[{"op": "replace", "path": "/spec/template/spec/containers/0/env/0", "value": {"name": "WATCH_NAMESPACES", "value": "$(DEPLOYMENT_WATCH_NAMESPACES)"}}]'

	# Generate namespace overlay with dns-operator and dns provider resources
	cd $(CLUSTER_OVERLAY_DIR)/$(CLUSTER_NAME)/namespace-$(DEPLOYMENT_NAMESPACE) && \
	touch kustomization.yaml && \
	$(KUSTOMIZE) edit set namespace $(DEPLOYMENT_NAMESPACE)  && \
	$(KUSTOMIZE) edit add resource "./dns-operator" && \
	$(KUSTOMIZE) edit add resource "../dns-providers"
