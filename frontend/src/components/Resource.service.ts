import http from '@/helpers/http'

export default class ResourceService {
    public saveResource(namespaceName: string, resourceType: string, resourceName: string, labels: Map<string, string>, data: Map<string, string>) {
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
                return Promise.reject(`Unknown resource type: ${resourceType}`);
        }

        return fetch(`${this.baseUrl(resourceType)}/${namespaceName}/${resourceName}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(body),
        }).then(http.handleResponse);
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
        }).then(http.handleResponse);
    };

    public deleteResource(namespaceName: string, resourceType: string, resourceName: string) {
        return fetch(`${this.baseUrl(resourceType)}/${namespaceName}/${resourceName}`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
        }).then(http.handleResponse);
    };

    public getResource(namespaceName: string, resourceType: string, resourceName: string) {
        return fetch(
            `${this.baseUrl(resourceType)}/${namespaceName}/${resourceName}`
        ).then(http.handleResponse)
            .then((data: any) => {
                if (resourceType === "secret") {
                    if (data.data) {
                        // loop over the data and convert from base64 to string
                        Object.keys(data.data).forEach((key) => {
                            data.data[key] = atob(data.data[key]);
                        });
                    } else {
                        data.data = {};
                    }
                }
                return data;
            });
    };

    public getResources(namespaceName: string, resourceType: string) {
        return fetch(
            `${this.baseUrl(resourceType)}/${namespaceName}`
        ).then(http.handleResponse);
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
