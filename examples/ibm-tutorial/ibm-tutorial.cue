package ssf

secret: "kube-api-secret": {
	metadata: annotations: "kubernetes.io/service-account.name": "pipeline-account"
	type: "kubernetes.io/service-account-token"
}

serviceAccount: "pipeline-account": {
}

role: "pipeline-role": rules: [{
	apiGroups: [""]
	resources: ["services"]
	verbs: ["get", "create", "update", "patch"]
}, {
	apiGroups: ["apps"]
	resources: ["deployments"]
	verbs: ["get", "create", "update", "patch"]
}]

roleBinding: "pipeline-role-binding": {
	roleRef: {
		apiGroup: "rbac.authorization.k8s.io"
		kind:     "Role"
		name:     "pipeline-role"
	}
	subjects: [{
		kind: "ServiceAccount"
		name: "pipeline-account"
	}]
}

pipeline: "build-and-deploy-pipeline": spec: {
	workspaces: [{
		name:        "git-source"
		description: "The git repo"
	}]
	params: [{
		name:        "gitUrl"
		description: "Git repository url"
	}, {
		name:        "gitRevision"
		description: "Git revision to check out"
		default:     "master"
	}, {
		name:        "pathToContext"
		description: "The path to the build context, used by Kaniko - within the workspace"
		default:     "src"
	}, {
		name:        "pathToYamlFile"
		description: "The path to the yaml file to deploy within the git source"
	}, {
		name:        "imageUrl"
		description: "Image name including repository"
	}, {
		name:        "imageTag"
		description: "Image tag"
		default:     "latest"
	}]
	tasks: [{
		name: "clone-repo"
		taskRef: name: "git-clone"
		workspaces: [{
			name:      "output"
			workspace: "git-source"
		}]
		params: [{
			name:  "url"
			value: "$(params.gitUrl)"
		}, {
			name:  "revision"
			value: "$(params.gitRevision)"
		}, {
			name:  "subdirectory"
			value: "."
		}, {
			name:  "deleteExisting"
			value: "true"
		}]
	}, {
		name: "source-to-image"
		taskRef: name: "kaniko-chains"
		runAfter: [
			"clone-repo",
		]
		workspaces: [{
			name:      "source"
			workspace: "git-source"
		}]
		params: [{
			name:  "CONTEXT"
			value: "$(params.pathToContext)"
		}, {
			name:  "IMAGE"
			value: "$(params.imageUrl):$(params.imageTag)"
		}]
	}]
}

pipelineRun: "picalc-pr-": spec: {
	pipelineRef: name: "build-and-deploy-pipeline"
	params: [{
		name:  "gitUrl"
		value: "https://github.com/IBM/tekton-tutorial"
	}, {
		name:  "gitRevision"
		value: "beta-update"
	}, {
		name:  "pathToYamlFile"
		value: "kubernetes/picalc.yaml"
	}, {
		name:  "imageUrl"
		value: "registry.kube-system.svc.cluster.local/citi/picalc"
	}, {
		name:  "imageTag"
		value: "1.0"
	}]
	serviceAccountName: "pipeline-account"
	workspaces: [{
		name: "git-source"
	}]
}