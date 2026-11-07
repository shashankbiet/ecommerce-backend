const authService = require("../../../app/auth/auth-service");
const authController = require("../../../app/auth/auth-controller");
const ApiError = require("../../../app/shared/models/api-error");

jest.mock("../../../app/auth/auth-service", () => ({
    authenticate: jest.fn(),
}));

const makeRes = () => {
    const res = {};
    res.status = jest.fn().mockReturnValue(res);
    res.send = jest.fn().mockReturnValue(res);
    return res;
};

describe("authController.authenticate", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    it("returns 200 with the user when credentials are valid", async () => {
        const user = { userId: 1, email: "a@b.com" };
        authService.authenticate.mockResolvedValue(user);
        const res = makeRes();

        await authController.authenticate(
            { body: { email: "a@b.com", password: "pw" } },
            res,
            jest.fn()
        );

        expect(authService.authenticate).toHaveBeenCalledWith(
            "a@b.com",
            undefined,
            "pw"
        );
        expect(res.status).toHaveBeenCalledWith(200);
        expect(res.send).toHaveBeenCalledWith(user);
    });

    it("returns 401 when the service finds no matching user", async () => {
        authService.authenticate.mockResolvedValue(null);
        const res = makeRes();

        await authController.authenticate(
            { body: { phoneNumber: "123", password: "bad" } },
            res,
            jest.fn()
        );

        expect(res.status).toHaveBeenCalledWith(401);
    });

    it("passes service errors to next as a 500 ApiError", async () => {
        authService.authenticate.mockRejectedValue(new Error("db down"));
        const next = jest.fn();

        await authController.authenticate(
            { body: { email: "a@b.com", password: "pw" } },
            makeRes(),
            next
        );

        expect(next).toHaveBeenCalledWith(expect.any(ApiError));
        expect(next.mock.calls[0][0].status).toBe(500);
    });
});
