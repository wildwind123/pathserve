export interface Message {
    fromPathServe: true,
    key: string,
    scenario : 'setParentValue' | 'setChildValue'
    Data: Value
    Form: FormInput
}

export type Value = ValueString | ValueNumber

export interface ValueString {
    type: 'string',
    data : string
}

export interface ValueNumber {
    type: 'number',
    data: number
}

export interface FormInput {
    name: string
    element: 'input'
    type: 'text' | 'number'
}