export default {
    handleResponse(response: Response) {
        if (!response.ok) {
            return response.text().then((text: string) => {
                const error = response.statusText + ': ' + text
                return Promise.reject(error)
            })
        }
        return response.text().then((text: string) => {
            const data = text && JSON.parse(text)
            return data
        })
    },
    handlePlainTextResponse(response: Response) {
        if (!response.ok) {
            return response.text().then((text: string) => {
                const error = response.statusText + ': ' + text
                return Promise.reject(error)
            })
        }
        return response.text().then((text: string) => {
            return text
        })
    },
}
