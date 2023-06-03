const publicUrl = ["/api/user/register"];

module.exports = {
    isPublicUrl: (originalUrl) => publicUrl.indexOf(originalUrl) >= 0,
};
