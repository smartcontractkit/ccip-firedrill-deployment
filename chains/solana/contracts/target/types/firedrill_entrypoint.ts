export type FiredrillEntrypoint = {
  "version": "0.1.0",
  "name": "firedrill_entrypoint",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "entrypoint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "chainSelector",
          "type": "u64"
        },
        {
          "name": "token",
          "type": "publicKey"
        },
        {
          "name": "offRamp",
          "type": "publicKey"
        },
        {
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "compound",
          "type": "publicKey"
        },
        {
          "name": "receiver",
          "type": "publicKey"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "firedrillEntrypoint",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "token",
            "type": "publicKey"
          },
          {
            "name": "offRamp",
            "type": "publicKey"
          },
          {
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "compound",
            "type": "publicKey"
          },
          {
            "name": "receiver",
            "type": "publicKey"
          },
          {
            "name": "sendLast",
            "type": "u8"
          }
        ]
      }
    }
  ]
};

export const IDL: FiredrillEntrypoint = {
  "version": "0.1.0",
  "name": "firedrill_entrypoint",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "entrypoint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "authority",
          "isMut": true,
          "isSigner": true
        },
        {
          "name": "systemProgram",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "chainSelector",
          "type": "u64"
        },
        {
          "name": "token",
          "type": "publicKey"
        },
        {
          "name": "offRamp",
          "type": "publicKey"
        },
        {
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "compound",
          "type": "publicKey"
        },
        {
          "name": "receiver",
          "type": "publicKey"
        }
      ]
    }
  ],
  "accounts": [
    {
      "name": "firedrillEntrypoint",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "token",
            "type": "publicKey"
          },
          {
            "name": "offRamp",
            "type": "publicKey"
          },
          {
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "compound",
            "type": "publicKey"
          },
          {
            "name": "receiver",
            "type": "publicKey"
          },
          {
            "name": "sendLast",
            "type": "u8"
          }
        ]
      }
    }
  ]
};
