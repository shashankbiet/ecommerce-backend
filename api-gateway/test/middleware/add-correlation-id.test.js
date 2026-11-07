jest.mock("uuid", () => ({
    v4: jest.fn(() => "generated-correlation-id"),
}));

jest.mock("../../handler/logger", () => ({
    error: jest.fn(),
}));

const addCorrelationId = require("../../middleware/add-correlation-id");
const apiStatus = require("../../app/shared/constants/api-status");

const makeRes = () => {
    const res = {};
    res.setHeader = jest.fn();
    res.status = jest.fn().mockReturnValue(res);
    res.send = jest.fn().mockReturnValue(res);
    return res;
};

describe("addCorrelationId", () => {
    it("uses incoming correlation-id when present", () => {
        const req = {
            headers: {
                "correlation-id": "incoming-id",
            },
        };
        const res = makeRes();
        const next = jest.fn();

        addCorrelationId(req, res, next);

        expect(req.headers["correlation-id"]).toBe("incoming-id");
        expect(res.setHeader).toHaveBeenCalledWith(
            "correlation-id",
            "incoming-id"
        );
        expect(next).toHaveBeenCalledTimes(1);
    });

    it("generates correlation-id when missing", () => {
        const req = { headers: {} };
        const res = makeRes();
        const next = jest.fn();

        addCorrelationId(req, res, next);

        expect(req.headers["correlation-id"]).toBe("generated-correlation-id");
        expect(res.setHeader).toHaveBeenCalledWith(
            "correlation-id",
            "generated-correlation-id"
        );
        expect(next).toHaveBeenCalledTimes(1);
    });

    it("returns 500 when middleware throws", () => {
        const req = { headers: {} };
        const res = makeRes();
        const next = jest.fn();
        res.setHeader.mockImplementation(() => {
            throw new Error("write failed");
        });

        addCorrelationId(req, res, next);

        expect(next).not.toHaveBeenCalled();
        expect(res.status).toHaveBeenCalledWith(
            apiStatus.INTERNAL_SERVER_ERROR.status
        );
        expect(res.send).toHaveBeenCalledWith({
            success: false,
            result: apiStatus.INTERNAL_SERVER_ERROR.result,
        });
    });
});
