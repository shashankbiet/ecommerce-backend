const publicUrls = ["/api/user/register"];

module.exports = {
    isPublicUrl: (originalUrl) => {
        if (!originalUrl) return false;
        const path = originalUrl.split("?")[0].replace(/\/+$/, "");
        return publicUrls.includes(path);
    },
};
