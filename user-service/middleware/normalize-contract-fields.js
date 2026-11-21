const toSnakeCase = (value) =>
    value
        .replace(/([a-z0-9])([A-Z])/g, "$1_$2")
        .replace(/([A-Z])([A-Z][a-z])/g, "$1_$2")
        .toLowerCase();

const normalizeValue = (value) => {
    if (Array.isArray(value)) {
        return value.map(normalizeValue);
    }
    if (value && typeof value === "object") {
        return Object.fromEntries(
            Object.entries(value).map(([key, child]) => [
                toSnakeCase(key),
                normalizeValue(child),
            ])
        );
    }
    return value;
};

const normalizeContractFields = (req, res, next) => {
    req.body = normalizeValue(req.body);
    next();
};

module.exports = normalizeContractFields;
