export default class NamespaceService {

    public getNamespaces() {
        return fetch(
            `/api/namespaces`
        ).then((response) => response.json());
    };
}
