
export default class ResourceService {
    public saveResource(namespaceName: string, resourceType: string, resourceName: string, labels: Map<string, string>, data: Map<string, string>) {
        console.log("saveResource", namespaceName, resourceName, labels, data);
        const body: any = {
            "metadata": {
                "name": resourceName,
                "namespace": namespaceName,
                "labels": Object.fromEntries(labels),
            },
        };

        switch (resourceType) {
            case "configmap":
                body.data = Object.fromEntries(data);
                break;
            case "secret":
                body.stringData = Object.fromEntries(data);
                break;
            default:
                throw new Error(`Unknown resource type: ${resourceType}`);
        }

        return fetch(`${this.baseUrl(resourceType)}/${namespaceName}/${resourceName}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(body),
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
        ).then((response) => {
            return response.json().then((data: any) => {
                if (resourceType === "secret") {
                    // loop over the data and convert from base64 to string
                    Object.keys(data.data).forEach((key) => {
                        data.data[key] = atob(data.data[key]);
                    });
                }
                return data;
            })
        });
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
