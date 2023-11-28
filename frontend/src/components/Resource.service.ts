export default class ResourceService {
    public saveResource(namespaceName: string, resourceType: string, resourceName: string, labels: Map<string, string>, data: Map<string, string>) {
        console.log("saveResource", namespaceName, resourceName, labels, data);
        return fetch(`${this.baseUrl(resourceType)}/${namespaceName}/${resourceName}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                "metadata": {
                    "name": resourceName,
                    "namespace": namespaceName,
                    "labels": Object.fromEntries(labels),
                },
                "data": Object.fromEntries(data),
            }),
        });
    };

    public createResource(namespaceName: string, resourceType: string, resourceName: string) {
        return fetch(`${this.baseUrl(resourceType)}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                "metadata": {
                    "name": resourceName,
                    "namespace": namespaceName,
                },
            }),
        });
    };

    public deleteResource(namespaceName: string, resourceType: string, resourceName: string) {
        return fetch(`${this.baseUrl(resourceType)}/${namespaceName}/${resourceName}`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
        });
    };

    public getResource(namespaceName: string, resourceType: string, resourceName: string) {
        return fetch(
            `${this.baseUrl(resourceType)}/${namespaceName}/${resourceName}`
        ).then((response) => response.json());
    };

    public getResources(namespaceName: string, resourceType: string) {
        return fetch(
            `${this.baseUrl(resourceType)}/${namespaceName}`
        ).then((response) => response.json());
    };

    private baseUrl(resourceType: string) {
        const prefix = "/api";
        switch (resourceType) {
            case "configmap":
                return `${prefix}/configs`;
            case "secret":
                return `${prefix}/secrets`;
            default:
                throw new Error(`Unknown resource type: ${resourceType}`);
        }
    }
}
