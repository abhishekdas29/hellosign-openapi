import { AttributeTypeMap } from "./models";
export declare class SubUnclaimedDraftTemplateSigner {
    "role": string;
    "name": string;
    "emailAddress": string;
    static discriminator: string | undefined;
    static attributeTypeMap: AttributeTypeMap;
    static getAttributeTypeMap(): AttributeTypeMap;
}