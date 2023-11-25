export default class ConfigMapService {
    public saveConfigMap(namespaceName: string, configMapName: string, labels: Map<string, string>, content: Map<string, string>) {
        console.log("saveConfigMap", namespaceName, configMapName, labels, content);
        return fetch(`/api/configs/${namespaceName}/${configMapName}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                "metadata": {
                    "name": configMapName,
                    "namespace": namespaceName,
                    "labels": Object.fromEntries(labels),
                },
                "data": content,
            }),
        });
    };

    public createConfigMap(namespaceName: string, configMapName: string) {
        return fetch(`/api/configs`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                "metadata": {
                    "name": configMapName,
                    "namespace": namespaceName,
                },
            }),
        });
    };

    public deleteConfigMap(namespaceName: string, configMapName: string) {
        return fetch(`/api/configs/${namespaceName}/${configMapName}`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
        });
    };

    public getConfigMap(namespaceName: string, configMapName: string) {
        return fetch(
            `/api/configs/${namespaceName}/${configMapName}`
        ).then((response) => response.json());
    };
}
