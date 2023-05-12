const getFileName = (path: string) => {
    return path.split("/").pop() ?? "unknown"
}

export default {
    getFileName,
}