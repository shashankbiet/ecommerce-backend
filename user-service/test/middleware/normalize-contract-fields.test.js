const normalizeContractFields = require("../../middleware/normalize-contract-fields");

describe("normalizeContractFields", () => {
    it("converts legacy camelCase request fields to snake_case", () => {
        const req = {
            body: {
                phoneNumber: "1234567890",
                profile: { firstName: "Ada", lastName: "Lovelace" },
            },
        };
        const next = jest.fn();

        normalizeContractFields(req, {}, next);

        expect(req.body).toEqual({
            phone_number: "1234567890",
            profile: { first_name: "Ada", last_name: "Lovelace" },
        });
        expect(next).toHaveBeenCalledTimes(1);
    });
});
