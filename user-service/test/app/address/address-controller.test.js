const addressService = require("../../../app/address/address-service");
const addressController = require("../../../app/address/address-controller");
const ApiError = require("../../../app/shared/models/api-error");

jest.mock("../../../app/address/address-service", () => ({
    add: jest.fn(),
    get: jest.fn(),
    update: jest.fn(),
    delete: jest.fn(),
}));

const makeRes = () => {
    const res = {};
    res.status = jest.fn().mockReturnValue(res);
    res.send = jest.fn().mockReturnValue(res);
    return res;
};

const makeReq = (overrides = {}) => ({
    body: {},
    params: {},
    headers: { "user-id": "5" },
    ...overrides,
});

describe("addressController", () => {
    beforeEach(() => {
        jest.clearAllMocks();
    });

    it("add returns 201 when the address is saved", async () => {
        addressService.add.mockResolvedValue(true);
        const res = makeRes();

        await addressController.add(
            makeReq({ body: { line1: "x" } }),
            res,
            jest.fn()
        );

        expect(addressService.add).toHaveBeenCalledWith(5, { line1: "x" });
        expect(res.status).toHaveBeenCalledWith(201);
    });

    it("add forwards a 404 ApiError when the user is not found", async () => {
        addressService.add.mockResolvedValue(false);
        const next = jest.fn();

        await addressController.add(makeReq(), makeRes(), next);

        expect(next.mock.calls[0][0]).toBeInstanceOf(ApiError);
        expect(next.mock.calls[0][0].status).toBe(404);
    });

    it("getById returns 200 with the address", async () => {
        addressService.get.mockResolvedValue({ line1: "x" });
        const res = makeRes();

        await addressController.getById(
            makeReq({ params: { addressId: "a1" } }),
            res,
            jest.fn()
        );

        expect(addressService.get).toHaveBeenCalledWith(5, "a1");
        expect(res.status).toHaveBeenCalledWith(200);
    });

    it("get returns 200 with all addresses", async () => {
        addressService.get.mockResolvedValue([{ line1: "x" }]);
        const res = makeRes();

        await addressController.get(makeReq(), res, jest.fn());

        expect(addressService.get).toHaveBeenCalledWith(5);
        expect(res.status).toHaveBeenCalledWith(200);
    });

    it("update returns 200 when the address is updated", async () => {
        addressService.update.mockResolvedValue(true);
        const res = makeRes();

        await addressController.update(
            makeReq({ params: { addressId: "a1" }, body: { line1: "y" } }),
            res,
            jest.fn()
        );

        expect(addressService.update).toHaveBeenCalledWith(5, "a1", {
            line1: "y",
        });
        expect(res.status).toHaveBeenCalledWith(200);
    });

    it("delete returns 200 when the address is removed", async () => {
        addressService.delete.mockResolvedValue(true);
        const res = makeRes();

        await addressController.delete(
            makeReq({ params: { addressId: "a1" } }),
            res,
            jest.fn()
        );

        expect(addressService.delete).toHaveBeenCalledWith(5, "a1");
        expect(res.status).toHaveBeenCalledWith(200);
    });

    it("passes service errors to next as a 500 ApiError", async () => {
        addressService.delete.mockRejectedValue(new Error("db down"));
        const next = jest.fn();

        await addressController.delete(
            makeReq({ params: { addressId: "a1" } }),
            makeRes(),
            next
        );

        expect(next.mock.calls[0][0].status).toBe(500);
    });
});
