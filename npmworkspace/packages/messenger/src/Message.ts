export interface Message {
    fromPathServe: true,
    name: string,
    scenario : 'setControlValue' | 'setComponentValue'
    Data: Value | null
    Form: FormInput | FormControl
}

export type Value = ValueString | ValueNumber | ValueObject

export interface ValueString {
    type: 'string',
    data : string
}

export interface ValueNumber {
    type: 'number',
    data: number
}
export interface ValueObject {
    type: 'object',
    data: object
}

export interface FormInput {
    element: 'input'
    type: 'text' | 'number' | 'textarea'
}

export interface FormControl {
    element: 'button'
}