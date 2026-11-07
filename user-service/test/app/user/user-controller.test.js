const userService = require("../../../app/user/user-service");
const userController = require("../../../app/user/user-controller");
const ApiError = require("../../../app/shared/models/api-error");

jest.mock("../../../app/user/user-service", () => ({
    register: jest.fn(),
    isAdmin: jest.fn(),
    get: jest.fn(),
    update: jest.fn(),
    delete: jest.fn(),
    updatePassword: jest.fn(),
}));

const makeRes = () => {
    const res = {};
    res.status = jest.fn().mockReturnValue(res);
    res.send = jest.fn().mockReturnValue(res);
    return res;
};

const makeReq = (overrides = {}) => ({
    body: {},
    headers: { "user-id": "5" },
    ...overrides,
});

describe("userController", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    describe("register", () => {
        it("returns 201 when the user is created", async () => {
            userService.register.mockResolvedValue(true);
            const res = makeRes();

            await userController.register(
                makeReq({ body: { role: "CUSTOMER" } }),
                res,
                jest.fn()
            );

            expect(res.status).toHaveBeenCalledWith(201);
        });

        it("returns 409 when the user already exists", async () => {
            userService.register.mockResolvedValue(false);
            const res = makeRes();

            await userController.register(
                makeReq({ body: { role: "CUSTOMER" } }),
                res,
                jest.fn()
            );

            expect(res.status).toHaveBeenCalledWith(409);
        });

        it("returns 401 when an ADMIN is requested without a valid admin token", async () => {
            userService.isAdmin.mockReturnValue(false);
            const res = makeRes();

            await userController.register(
                makeReq({ body: { role: "ADMIN", token: "bad" } }),
                res,
                jest.fn()
            );

            expect(res.status).toHaveBeenCalledWith(401);
            expect(userService.register).not.toHaveBeenCalled();
        });

        it("passes service errors to next as a 500 ApiError", async () => {
            userService.register.mockRejectedValue(new Error("db down"));
            const next = jest.fn();

            await userController.register(
                makeReq({ body: { role: "CUSTOMER" } }),
                makeRes(),
                next
            );

            expect(next).toHaveBeenCalledWith(expect.any(ApiError));
            expect(next.mock.calls[0][0].status).toBe(500);
        });
    });

    describe("get", () => {
        it("looks up the user from the user-id header", async () => {
            userService.get.mockResolvedValue({ userId: 5 });
            const res = makeRes();

            await userController.get(makeReq(), res, jest.fn());

            expect(userService.get).toHaveBeenCalledWith(5);
            expect(res.status).toHaveBeenCalledWith(200);
            expect(res.send).toHaveBeenCalledWith({ userId: 5 });
        });

        it("returns 404 when the user does not exist", async () => {
            userService.get.mockResolvedValue(null);
            const res = makeRes();

            await userController.get(makeReq(), res, jest.fn());

            expect(res.status).toHaveBeenCalledWith(404);
        });
    });

    describe("update", () => {
        it("returns 200 when the update succeeds", async () => {
            userService.update.mockResolvedValue(true);
            const res = makeRes();

            await userController.update(
                makeReq({ body: { firstName: "A" } }),
                res,
                jest.fn()
            );

            expect(userService.update).toHaveBeenCalledWith(5, {
                firstName: "A",
            });
            expect(res.status).toHaveBeenCalledWith(200);
        });

        it("returns 404 when no user was updated", async () => {
            userService.update.mockResolvedValue(false);
            const res = makeRes();

            await userController.update(makeReq(), res, jest.fn());

            expect(res.status).toHaveBeenCalledWith(404);
        });
    });

    describe("delete", () => {
        it("returns 200 when the user is deleted", async () => {
            userService.delete.mockResolvedValue(true);
            const res = makeRes();

            await userController.delete(makeReq(), res, jest.fn());

            expect(userService.delete).toHaveBeenCalledWith(5);
            expect(res.status).toHaveBeenCalledWith(200);
        });

        it("returns 404 when the user is not found", async () => {
            userService.delete.mockResolvedValue(false);
            const res = makeRes();

            await userController.delete(makeReq(), res, jest.fn());

            expect(res.status).toHaveBeenCalledWith(404);
        });
    });

    describe("updatePassword", () => {
        it("rejects a new password equal to the current one", async () => {
            const res = makeRes();

            await userController.updatePassword(
                makeReq({
                    body: { currentPassword: "same", newPassword: "same" },
                }),
                res,
                jest.fn()
            );

            expect(res.status).toHaveBeenCalledWith(400);
            expect(userService.updatePassword).not.toHaveBeenCalled();
        });

        it("returns 200 when the password is changed", async () => {
            userService.updatePassword.mockResolvedValue(true);
            const res = makeRes();

            await userController.updatePassword(
                makeReq({
                    body: { currentPassword: "old", newPassword: "new" },
                }),
                res,
                jest.fn()
            );

            expect(res.status).toHaveBeenCalledWith(200);
        });

        it("returns 404 when the user is not found", async () => {
            userService.updatePassword.mockResolvedValue(null);
            const res = makeRes();

            await userController.updatePassword(
                makeReq({
                    body: { currentPassword: "old", newPassword: "new" },
                }),
                res,
                jest.fn()
            );

            expect(res.status).toHaveBeenCalledWith(404);
        });

        it("returns 400 when the current password is incorrect", async () => {
            userService.updatePassword.mockResolvedValue(false);
            const res = makeRes();

            await userController.updatePassword(
                makeReq({
                    body: { currentPassword: "old", newPassword: "new" },
                }),
                res,
                jest.fn()
            );

            expect(res.status).toHaveBeenCalledWith(400);
            expect(res.send).toHaveBeenCalledWith({
                success: false,
                result: "Current Password Incorrect",
            });
        });
    });
});
