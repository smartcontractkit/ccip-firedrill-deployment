export type FailingReceiver = {
  "version": "0.1.0",
  "name": "failing_receiver",
  "instructions": [
    {
      "name": "ccipReceive",
      "accounts": [],
      "args": [
        {
          "name": "message",
          "type": {
            "defined": "Any2SVMMessage"
          }
        }
      ]
    }
  ],
  "types": [
    {
      "name": "Any2SVMMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "messageId",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "sender",
            "type": "bytes"
          },
          {
            "name": "data",
            "type": "bytes"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": "SVMTokenAmount"
              }
            }
          }
        ]
      }
    },
    {
      "name": "SVMTokenAmount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "publicKey"
          },
          {
            "name": "amount",
            "type": "u64"
          }
        ]
      }
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "AlwaysErrors",
      "msg": "Always errors."
    }
  ]
};

export const IDL: FailingReceiver = {
  "version": "0.1.0",
  "name": "failing_receiver",
  "instructions": [
    {
      "name": "ccipReceive",
      "accounts": [],
      "args": [
        {
          "name": "message",
          "type": {
            "defined": "Any2SVMMessage"
          }
        }
      ]
    }
  ],
  "types": [
    {
      "name": "Any2SVMMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "messageId",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "sender",
            "type": "bytes"
          },
          {
            "name": "data",
            "type": "bytes"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": "SVMTokenAmount"
              }
            }
          }
        ]
      }
    },
    {
      "name": "SVMTokenAmount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "token",
            "type": "publicKey"
          },
          {
            "name": "amount",
            "type": "u64"
          }
        ]
      }
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "AlwaysErrors",
      "msg": "Always errors."
    }
  ]
};
