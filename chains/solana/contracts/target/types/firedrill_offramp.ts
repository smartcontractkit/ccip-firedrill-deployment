export type FiredrillOfframp = {
  "version": "0.1.0",
  "name": "firedrill_offramp",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "offramp",
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
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "compound",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "initializeConfig",
      "docs": [
        "Initializes the CCIP Offramp Config account.",
        "",
        "The initialization of the Offramp is responsibility of Admin, nothing more than calling these",
        "initialization methods should be done first.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for initialization of the config.",
        "* `svm_chain_selector` - The chain selector for SVM.",
        "* `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed."
      ],
      "accounts": [
        {
          "name": "config",
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
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "svmChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "emitSourceChainAdded",
      "accounts": [
        {
          "name": "offramp",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "emitCommitReportAccepted",
      "accounts": [
        {
          "name": "offramp",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "The OffRamp account—for ownership verification."
          ]
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        }
      ],
      "args": [
        {
          "name": "minSeqNr",
          "type": "u64"
        },
        {
          "name": "maxSeqNr",
          "type": "u64"
        }
      ]
    }
  ],
  "accounts": [
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
            "type": "u8"
          },
          {
            "name": "padding0",
            "type": {
              "array": [
                "u8",
                6
              ]
            }
          },
          {
            "name": "svmChainSelector",
            "type": "u64"
          },
          {
            "name": "enableManualExecutionAfter",
            "type": "i64"
          },
          {
            "name": "padding1",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
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
            "name": "padding2",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
          },
          {
            "name": "ocr3",
            "type": {
              "array": [
                {
                  "defined": "Ocr3Config"
                },
                2
              ]
            }
          }
        ]
      }
    },
    {
      "name": "referenceAddresses",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "router",
            "type": "publicKey"
          },
          {
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "offrampLookupTable",
            "type": "publicKey"
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "sourceChain",
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
              "defined": "SourceChainState"
            }
          },
          {
            "name": "config",
            "type": {
              "defined": "SourceChainConfig"
            }
          }
        ]
      }
    },
    {
      "name": "commitReport",
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
            "name": "merkleRoot",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "timestamp",
            "type": "i64"
          },
          {
            "name": "minMsgNr",
            "type": "u64"
          },
          {
            "name": "maxMsgNr",
            "type": "u64"
          },
          {
            "name": "executionStates",
            "type": "u128"
          }
        ]
      }
    },
    {
      "name": "firedrillOffRamp",
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
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "compound",
            "type": "publicKey"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "PriceUpdates",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenPriceUpdates",
            "type": {
              "vec": {
                "defined": "TokenPriceUpdate"
              }
            }
          },
          {
            "name": "gasPriceUpdates",
            "type": {
              "vec": {
                "defined": "GasPriceUpdate"
              }
            }
          }
        ]
      }
    },
    {
      "name": "TokenPriceUpdate",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceToken",
            "type": "publicKey"
          },
          {
            "name": "usdPerToken",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          }
        ]
      }
    },
    {
      "name": "GasPriceUpdate",
      "docs": [
        "Gas price for a given chain in USD; its value may contain tightly packed fields."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "usdPerUnitGas",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          }
        ]
      }
    },
    {
      "name": "MerkleRoot",
      "docs": [
        "Struct to hold a merkle root and an interval for a source chain"
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "onRampAddress",
            "type": "bytes"
          },
          {
            "name": "minSeqNr",
            "type": "u64"
          },
          {
            "name": "maxSeqNr",
            "type": "u64"
          },
          {
            "name": "merkleRoot",
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
      "name": "ConfigOcrPluginType",
      "docs": [
        "It's not possible to store enums in zero_copy accounts, so we wrap the discriminant",
        "in a struct to store in config."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "discriminant",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "Ocr3ConfigInfo",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "configDigest",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "f",
            "type": "u8"
          },
          {
            "name": "n",
            "type": "u8"
          },
          {
            "name": "isSignatureVerificationEnabled",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "Ocr3Config",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "pluginType",
            "type": {
              "defined": "ConfigOcrPluginType"
            }
          },
          {
            "name": "configInfo",
            "type": {
              "defined": "Ocr3ConfigInfo"
            }
          },
          {
            "name": "signers",
            "type": {
              "array": [
                {
                  "array": [
                    "u8",
                    20
                  ]
                },
                16
              ]
            }
          },
          {
            "name": "transmitters",
            "type": {
              "array": [
                {
                  "array": [
                    "u8",
                    32
                  ]
                },
                16
              ]
            }
          }
        ]
      }
    },
    {
      "name": "SourceChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "isEnabled",
            "type": "bool"
          },
          {
            "name": "isRmnVerificationDisabled",
            "type": "bool"
          },
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "onRamp",
            "type": {
              "defined": "OnRampAddress"
            }
          }
        ]
      }
    },
    {
      "name": "OnRampAddress",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "bytes",
            "type": {
              "array": [
                "u8",
                64
              ]
            }
          },
          {
            "name": "len",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "SourceChainState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "minSeqNr",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "OcrPluginType",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Commit"
          },
          {
            "name": "Execution"
          }
        ]
      }
    },
    {
      "name": "MessageExecutionState",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Untouched"
          },
          {
            "name": "InProgress"
          },
          {
            "name": "Success"
          },
          {
            "name": "Failure"
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
    }
  ],
  "events": [
    {
      "name": "CommitReportAccepted",
      "fields": [
        {
          "name": "merkleRoot",
          "type": {
            "option": {
              "defined": "MerkleRoot"
            }
          },
          "index": false
        },
        {
          "name": "priceUpdates",
          "type": {
            "defined": "PriceUpdates"
          },
          "index": false
        }
      ]
    },
    {
      "name": "ExecutionStateChanged",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "messageId",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "messageHash",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "state",
          "type": {
            "defined": "MessageExecutionState"
          },
          "index": false
        }
      ]
    },
    {
      "name": "SourceChainAdded",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "ConfigSet",
      "fields": [
        {
          "name": "svmChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "enableManualExecutionAfter",
          "type": "i64",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidOnrampAddress",
      "msg": "Invalid OnRamp address"
    },
    {
      "code": 6001,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    },
    {
      "code": 6002,
      "name": "InvalidVersion",
      "msg": "Invalid version"
    },
    {
      "code": 6003,
      "name": "InvalidPluginType",
      "msg": "Invalid plugin type"
    },
    {
      "code": 6004,
      "name": "Unauthorized",
      "msg": "Unauthorized"
    }
  ]
};

export const IDL: FiredrillOfframp = {
  "version": "0.1.0",
  "name": "firedrill_offramp",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "referenceAddresses",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "offramp",
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
          "name": "feeQuoter",
          "type": "publicKey"
        },
        {
          "name": "compound",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "initializeConfig",
      "docs": [
        "Initializes the CCIP Offramp Config account.",
        "",
        "The initialization of the Offramp is responsibility of Admin, nothing more than calling these",
        "initialization methods should be done first.",
        "",
        "# Arguments",
        "",
        "* `ctx` - The context containing the accounts required for initialization of the config.",
        "* `svm_chain_selector` - The chain selector for SVM.",
        "* `enable_execution_after` - The minimum amount of time required between a message has been committed and can be manually executed."
      ],
      "accounts": [
        {
          "name": "config",
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
        },
        {
          "name": "program",
          "isMut": false,
          "isSigner": false
        },
        {
          "name": "programData",
          "isMut": false,
          "isSigner": false
        }
      ],
      "args": [
        {
          "name": "svmChainSelector",
          "type": "u64"
        }
      ]
    },
    {
      "name": "emitSourceChainAdded",
      "accounts": [
        {
          "name": "offramp",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        },
        {
          "name": "sourceChain",
          "isMut": true,
          "isSigner": false
        }
      ],
      "args": []
    },
    {
      "name": "emitCommitReportAccepted",
      "accounts": [
        {
          "name": "offramp",
          "isMut": true,
          "isSigner": false,
          "docs": [
            "The OffRamp account—for ownership verification."
          ]
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        }
      ],
      "args": [
        {
          "name": "minSeqNr",
          "type": "u64"
        },
        {
          "name": "maxSeqNr",
          "type": "u64"
        }
      ]
    }
  ],
  "accounts": [
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
            "type": "u8"
          },
          {
            "name": "padding0",
            "type": {
              "array": [
                "u8",
                6
              ]
            }
          },
          {
            "name": "svmChainSelector",
            "type": "u64"
          },
          {
            "name": "enableManualExecutionAfter",
            "type": "i64"
          },
          {
            "name": "padding1",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
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
            "name": "padding2",
            "type": {
              "array": [
                "u8",
                8
              ]
            }
          },
          {
            "name": "ocr3",
            "type": {
              "array": [
                {
                  "defined": "Ocr3Config"
                },
                2
              ]
            }
          }
        ]
      }
    },
    {
      "name": "referenceAddresses",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "version",
            "type": "u8"
          },
          {
            "name": "router",
            "type": "publicKey"
          },
          {
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "offrampLookupTable",
            "type": "publicKey"
          },
          {
            "name": "rmnRemote",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "sourceChain",
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
              "defined": "SourceChainState"
            }
          },
          {
            "name": "config",
            "type": {
              "defined": "SourceChainConfig"
            }
          }
        ]
      }
    },
    {
      "name": "commitReport",
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
            "name": "merkleRoot",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "timestamp",
            "type": "i64"
          },
          {
            "name": "minMsgNr",
            "type": "u64"
          },
          {
            "name": "maxMsgNr",
            "type": "u64"
          },
          {
            "name": "executionStates",
            "type": "u128"
          }
        ]
      }
    },
    {
      "name": "firedrillOffRamp",
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
            "name": "feeQuoter",
            "type": "publicKey"
          },
          {
            "name": "compound",
            "type": "publicKey"
          }
        ]
      }
    }
  ],
  "types": [
    {
      "name": "PriceUpdates",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "tokenPriceUpdates",
            "type": {
              "vec": {
                "defined": "TokenPriceUpdate"
              }
            }
          },
          {
            "name": "gasPriceUpdates",
            "type": {
              "vec": {
                "defined": "GasPriceUpdate"
              }
            }
          }
        ]
      }
    },
    {
      "name": "TokenPriceUpdate",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceToken",
            "type": "publicKey"
          },
          {
            "name": "usdPerToken",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          }
        ]
      }
    },
    {
      "name": "GasPriceUpdate",
      "docs": [
        "Gas price for a given chain in USD; its value may contain tightly packed fields."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "destChainSelector",
            "type": "u64"
          },
          {
            "name": "usdPerUnitGas",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          }
        ]
      }
    },
    {
      "name": "MerkleRoot",
      "docs": [
        "Struct to hold a merkle root and an interval for a source chain"
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "sourceChainSelector",
            "type": "u64"
          },
          {
            "name": "onRampAddress",
            "type": "bytes"
          },
          {
            "name": "minSeqNr",
            "type": "u64"
          },
          {
            "name": "maxSeqNr",
            "type": "u64"
          },
          {
            "name": "merkleRoot",
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
      "name": "ConfigOcrPluginType",
      "docs": [
        "It's not possible to store enums in zero_copy accounts, so we wrap the discriminant",
        "in a struct to store in config."
      ],
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "discriminant",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "Ocr3ConfigInfo",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "configDigest",
            "type": {
              "array": [
                "u8",
                32
              ]
            }
          },
          {
            "name": "f",
            "type": "u8"
          },
          {
            "name": "n",
            "type": "u8"
          },
          {
            "name": "isSignatureVerificationEnabled",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "Ocr3Config",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "pluginType",
            "type": {
              "defined": "ConfigOcrPluginType"
            }
          },
          {
            "name": "configInfo",
            "type": {
              "defined": "Ocr3ConfigInfo"
            }
          },
          {
            "name": "signers",
            "type": {
              "array": [
                {
                  "array": [
                    "u8",
                    20
                  ]
                },
                16
              ]
            }
          },
          {
            "name": "transmitters",
            "type": {
              "array": [
                {
                  "array": [
                    "u8",
                    32
                  ]
                },
                16
              ]
            }
          }
        ]
      }
    },
    {
      "name": "SourceChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "isEnabled",
            "type": "bool"
          },
          {
            "name": "isRmnVerificationDisabled",
            "type": "bool"
          },
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "onRamp",
            "type": {
              "defined": "OnRampAddress"
            }
          }
        ]
      }
    },
    {
      "name": "OnRampAddress",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "bytes",
            "type": {
              "array": [
                "u8",
                64
              ]
            }
          },
          {
            "name": "len",
            "type": "u32"
          }
        ]
      }
    },
    {
      "name": "SourceChainState",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "minSeqNr",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "OcrPluginType",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Commit"
          },
          {
            "name": "Execution"
          }
        ]
      }
    },
    {
      "name": "MessageExecutionState",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Untouched"
          },
          {
            "name": "InProgress"
          },
          {
            "name": "Success"
          },
          {
            "name": "Failure"
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
    }
  ],
  "events": [
    {
      "name": "CommitReportAccepted",
      "fields": [
        {
          "name": "merkleRoot",
          "type": {
            "option": {
              "defined": "MerkleRoot"
            }
          },
          "index": false
        },
        {
          "name": "priceUpdates",
          "type": {
            "defined": "PriceUpdates"
          },
          "index": false
        }
      ]
    },
    {
      "name": "ExecutionStateChanged",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sequenceNumber",
          "type": "u64",
          "index": false
        },
        {
          "name": "messageId",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "messageHash",
          "type": {
            "array": [
              "u8",
              32
            ]
          },
          "index": false
        },
        {
          "name": "state",
          "type": {
            "defined": "MessageExecutionState"
          },
          "index": false
        }
      ]
    },
    {
      "name": "SourceChainAdded",
      "fields": [
        {
          "name": "sourceChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "sourceChainConfig",
          "type": {
            "defined": "SourceChainConfig"
          },
          "index": false
        }
      ]
    },
    {
      "name": "ConfigSet",
      "fields": [
        {
          "name": "svmChainSelector",
          "type": "u64",
          "index": false
        },
        {
          "name": "enableManualExecutionAfter",
          "type": "i64",
          "index": false
        }
      ]
    }
  ],
  "errors": [
    {
      "code": 6000,
      "name": "InvalidOnrampAddress",
      "msg": "Invalid OnRamp address"
    },
    {
      "code": 6001,
      "name": "InvalidCodeVersion",
      "msg": "Invalid code version"
    },
    {
      "code": 6002,
      "name": "InvalidVersion",
      "msg": "Invalid version"
    },
    {
      "code": 6003,
      "name": "InvalidPluginType",
      "msg": "Invalid plugin type"
    },
    {
      "code": 6004,
      "name": "Unauthorized",
      "msg": "Unauthorized"
    }
  ]
};
