jest.mock("express-validator", () => ({
    validationResult: jest.fn(),
}));

const { validationResult } = require("express-validator");
const validate = require("../../util/validator");
const status = require("../../app/shared/constants/api-status");

const makeRes = () => {
    const res = {};
    res.status = jest.fn().mockReturnValue(res);
    res.json = jest.fn().mockReturnValue(res);
    return res;
};

describe("validate middleware", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    it("calls next when validation succeeds", () => {
        validationResult.mockReturnValue({
            isEmpty: () => true,
            array: () => [],
        });
        const req = {};
        const res = makeRes();
        const next = jest.fn();

        validate(req, res, next);

        expect(next).toHaveBeenCalledTimes(1);
        expect(res.status).not.toHaveBeenCalled();
    });

    it("returns mapped validation errors when validation fails", () => {
        validationResult.mockReturnValue({
            isEmpty: () => false,
            array: () => [
                { path: "email", msg: "Invalid email" },
                { param: "password", msg: "Required field" },
                { msg: "Unknown field error" },
            ],
        });
        const req = {};
        const res = makeRes();

        validate(req, res, jest.fn());

        expect(res.status).toHaveBeenCalledWith(status.BAD_REQUEST.status);
        expect(res.json).toHaveBeenCalledWith({
            errors: [
                { email: "Invalid email" },
                { password: "Required field" },
                { field: "Unknown field error" },
            ],
        });
    });
});
