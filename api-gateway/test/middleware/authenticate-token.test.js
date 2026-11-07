const jwt = require("jsonwebtoken");
const authenticateToken = require("../../middleware/authenticate-token");

jest.mock("../../config", () => ({ accessToken: "test-secret" }));

const makeRes = () => {
    const res = {};
    res.status = jest.fn().mockReturnValue(res);
    res.send = jest.fn().mockReturnValue(res);
    return res;
};

describe("authenticateToken", () => {
    let next;

    beforeEach(() => {
        next = jest.fn();
    });

    it("skips authentication for public URLs", () => {
        const req = { originalUrl: "/api/user/register", headers: {} };
        const res = makeRes();

        authenticateToken(req, res, next);

        expect(next).toHaveBeenCalledTimes(1);
        expect(res.status).not.toHaveBeenCalled();
    });

    it("returns 401 when the authorization header is missing", () => {
        const req = { originalUrl: "/api/user", headers: {} };
        const res = makeRes();

        authenticateToken(req, res, next);

        expect(res.status).toHaveBeenCalledWith(401);
        expect(res.send).toHaveBeenCalled();
        expect(next).not.toHaveBeenCalled();
    });

    it("accepts a valid Bearer token and forwards identity headers", () => {
        const token = jwt.sign({ userId: 7, role: "CUSTOMER" }, "test-secret");
        const req = {
            originalUrl: "/api/user",
            headers: { authorization: `Bearer ${token}` },
        };
        const res = makeRes();

        authenticateToken(req, res, next);

        expect(next).toHaveBeenCalledTimes(1);
        expect(req.headers["user-id"]).toBe(7);
        expect(req.headers["user-role"]).toBe("CUSTOMER");
        expect(res.status).not.toHaveBeenCalled();
    });

    it("accepts a raw token without the Bearer prefix", () => {
        const token = jwt.sign({ userId: 3, role: "ADMIN" }, "test-secret");
        const req = {
            originalUrl: "/api/user",
            headers: { authorization: token },
        };

        authenticateToken(req, makeRes(), next);

        expect(next).toHaveBeenCalledTimes(1);
        expect(req.headers["user-role"]).toBe("ADMIN");
    });

    it("returns 401 for a token signed with a different secret", () => {
        const token = jwt.sign({ userId: 1, role: "CUSTOMER" }, "wrong-secret");
        const req = {
            originalUrl: "/api/user",
            headers: { authorization: `Bearer ${token}` },
        };
        const res = makeRes();

        authenticateToken(req, res, next);

        expect(res.status).toHaveBeenCalledWith(401);
        expect(next).not.toHaveBeenCalled();
    });

    it("returns 401 for an expired token", () => {
        const expired = Math.floor(Date.now() / 1000) - 60;
        const token = jwt.sign(
            { userId: 1, role: "CUSTOMER", exp: expired },
            "test-secret"
        );
        const req = {
            originalUrl: "/api/user",
            headers: { authorization: `Bearer ${token}` },
        };
        const res = makeRes();

        authenticateToken(req, res, next);

        expect(res.status).toHaveBeenCalledWith(401);
        expect(next).not.toHaveBeenCalled();
    });
});
