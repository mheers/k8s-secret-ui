
const createConfigMap = (namespaceName: string, configMapName: string) => {
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

const deleteConfigMap = (namespaceName: string, configMapName: string) => {
    return fetch(`/api/configs/${namespaceName}/${configMapName}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
        },
    });
};

export { createConfigMap, deleteConfigMap };