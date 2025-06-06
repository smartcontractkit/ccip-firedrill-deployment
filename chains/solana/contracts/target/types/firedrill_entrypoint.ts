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
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "destChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "token",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "The PDA where we want a single “firedrill token” mint for this entrypoint"
          ]
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "SPL Token program"
          ]
        },
        {
          "name": "rent",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "Rent sysvar for SPL init"
          ]
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
          "name": "offRamp",
          "type": "publicKey"
        },
        {
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "receiver",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "emitCcipMessageSent",
      "accounts": [
        {
          "name": "entrypoint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        }
      ],
      "args": [
        {
          "name": "sender",
          "type": "publicKey"
        },
        {
          "name": "index",
          "type": "u64"
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
            "name": "receiver",
            "type": "publicKey"
          },
          {
            "name": "sendLast",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "destChain",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "state",
            "type": {
              "defined": "DestChainState"
            }
          },
          {
            "name": "config",
            "type": {
              "defined": "DestChainConfig"
            }
          }
        ]
      }
    },
    {
      "name": "config",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "defaultCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "svmChainSelector",
            "type": "u64"
          },
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          },
          {
            "name": "linkTokenMint",
            "type": "publicKey"
          },
          {
            "name": "feeAggregator",
            "type": "publicKey"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "RampMessageHeader",
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
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "sequenceNumber",
            "type": "u64"
          },
          {
            "name": "nonce",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "SVM2AnyRampMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "header",
            "type": {
              "defined": "RampMessageHeader"
            }
          },
          {
            "name": "sender",
            "type": "publicKey"
          },
          {
            "name": "data",
            "type": "bytes"
          },
          {
            "name": "receiver",
            "type": "bytes"
          },
          {
            "name": "extraArgs",
            "type": "bytes"
          },
          {
            "name": "feeToken",
            "type": "publicKey"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": "SVM2AnyTokenTransfer"
              }
            }
          },
          {
            "name": "feeTokenAmount",
            "type": {
              "defined": "CrossChainAmount"
            }
          },
          {
            "name": "feeValueJuels",
            "type": {
              "defined": "CrossChainAmount"
            }
          }
        ]
      }
    },
    {
      "name": "SVM2AnyTokenTransfer",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourcePoolAddress",
            "type": "publicKey"
          },
          {
            "name": "destTokenAddress",
            "type": "bytes"
          },
          {
            "name": "extraData",
            "type": "bytes"
          },
          {
            "name": "amount",
            "type": {
              "defined": "CrossChainAmount"
            }
          },
          {
            "name": "destExecData",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "CrossChainAmount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "leBytes",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          }
        ]
      }
    },
    {
      "name": "DestChainState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sequenceNumber",
            "type": "u64"
          },
          {
            "name": "sequenceNumberToRestore",
            "type": "u64"
          },
          {
            "name": "restoreOnAction",
            "type": {
              "defined": "RestoreOnAction"
            }
          }
        ]
      }
    },
    {
      "name": "DestChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "allowedSenders",
            "type": {
              "vec": "publicKey"
            }
          },
          {
            "name": "allowListEnabled",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "CodeVersion",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Default"
          },
          {
            "name": "V1"
          }
        ]
      }
    },
    {
      "name": "RestoreOnAction",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "None"
          },
          {
            "name": "Upgrade"
          },
          {
            "name": "Rollback"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "CCIPMessageSent",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "message",
          "type": {
            "defined": "SVM2AnyRampMessage"
          },
          "index": false
        }
      ]
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
          "name": "config",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "destChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "token",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "The PDA where we want a single “firedrill token” mint for this entrypoint"
          ]
        },
        {
          "name": "tokenProgram",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "SPL Token program"
          ]
        },
        {
          "name": "rent",
          "isMut": false,
          "isSigner": false,
          "docs": [
            "Rent sysvar for SPL init"
          ]
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
          "name": "offRamp",
          "type": "publicKey"
        },
        {
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "receiver",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "emitCcipMessageSent",
      "accounts": [
        {
          "name": "entrypoint",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        }
      ],
      "args": [
        {
          "name": "sender",
          "type": "publicKey"
        },
        {
          "name": "index",
          "type": "u64"
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
            "name": "receiver",
            "type": "publicKey"
          },
          {
            "name": "sendLast",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "destChain",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "state",
            "type": {
              "defined": "DestChainState"
            }
          },
          {
            "name": "config",
            "type": {
              "defined": "DestChainConfig"
            }
          }
        ]
      }
    },
    {
      "name": "config",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "defaultCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "svmChainSelector",
            "type": "u64"
          },
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          },
          {
            "name": "linkTokenMint",
            "type": "publicKey"
          },
          {
            "name": "feeAggregator",
            "type": "publicKey"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "RampMessageHeader",
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
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "sequenceNumber",
            "type": "u64"
          },
          {
            "name": "nonce",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "SVM2AnyRampMessage",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "header",
            "type": {
              "defined": "RampMessageHeader"
            }
          },
          {
            "name": "sender",
            "type": "publicKey"
          },
          {
            "name": "data",
            "type": "bytes"
          },
          {
            "name": "receiver",
            "type": "bytes"
          },
          {
            "name": "extraArgs",
            "type": "bytes"
          },
          {
            "name": "feeToken",
            "type": "publicKey"
          },
          {
            "name": "tokenAmounts",
            "type": {
              "vec": {
                "defined": "SVM2AnyTokenTransfer"
              }
            }
          },
          {
            "name": "feeTokenAmount",
            "type": {
              "defined": "CrossChainAmount"
            }
          },
          {
            "name": "feeValueJuels",
            "type": {
              "defined": "CrossChainAmount"
            }
          }
        ]
      }
    },
    {
      "name": "SVM2AnyTokenTransfer",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourcePoolAddress",
            "type": "publicKey"
          },
          {
            "name": "destTokenAddress",
            "type": "bytes"
          },
          {
            "name": "extraData",
            "type": "bytes"
          },
          {
            "name": "amount",
            "type": {
              "defined": "CrossChainAmount"
            }
          },
          {
            "name": "destExecData",
            "type": "bytes"
          }
        ]
      }
    },
    {
      "name": "CrossChainAmount",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "leBytes",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          }
        ]
      }
    },
    {
      "name": "DestChainState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sequenceNumber",
            "type": "u64"
          },
          {
            "name": "sequenceNumberToRestore",
            "type": "u64"
          },
          {
            "name": "restoreOnAction",
            "type": {
              "defined": "RestoreOnAction"
            }
          }
        ]
      }
    },
    {
      "name": "DestChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "allowedSenders",
            "type": {
              "vec": "publicKey"
            }
          },
          {
            "name": "allowListEnabled",
            "type": "bool"
          }
        ]
      }
    },
    {
      "name": "CodeVersion",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Default"
          },
          {
            "name": "V1"
          }
        ]
      }
    },
    {
      "name": "RestoreOnAction",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "None"
          },
          {
            "name": "Upgrade"
          },
          {
            "name": "Rollback"
          }
        ]
      }
    }
  ],
  "events": [
    {
      "name": "CCIPMessageSent",
      "fields": [
        {
          "name": "destChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "message",
          "type": {
            "defined": "SVM2AnyRampMessage"
          },
          "index": false
        }
      ]
    }
  ]
};
