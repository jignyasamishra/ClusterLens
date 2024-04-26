# ClusterLens
ClusterLens is a command-line tool for scanning Kubernetes clusters to generate a summary of resources.
# Features
- Scan Kubernetes Cluster: Quickly scan your Kubernetes cluster to gather essential information about pods.
- Namespace Support: Specify namespaces to focus the scan on specific parts of the cluster.
- Logging: Enable logging to track the scanning process and any encountered errors.
# Flags

- `-kubeconfig`: Specify the location of the kubeconfig file.
- `--logger`: Enable logging to track the scanning process.
- `--namespaces (-n)`: Provide a list of namespaces to focus the scan.
## Installation

Follow these steps to install ClusterLens:

1. **Prerequisites**: Ensure you have Go installed on your machine

2. **Clone the Repository**: Use the command `git clone https://github.com/jignyasamishra/clusterlens.git` to clone the repository.

3. **Navigate to the Directory**: Use the command `cd clusterlens` to navigate into the project directory.

4. **Build the Project**: Use the command `go build -o clusterlens .` to build the project. This command builds the project and outputs the binary file as `clusterlens`.
