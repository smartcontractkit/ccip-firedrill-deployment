export type FiredrillFeequoter = {
  "version": "0.1.0",
  "name": "firedrill_feequoter",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "feeQuoter",
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
          "isMut": false,
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
          "name": "compound",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "emitUsdPerTokenUpdated",
      "accounts": [
        {
          "name": "feeQuoter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        }
      ],
      "args": []
    }
  ],
  "accounts": [
    {
      "name": "firedrillFeeQuoter",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "compound",
            "type": "publicKey"
          },
          {
            "name": "token",
            "type": "publicKey"
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
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "maxFeeJuelsPerMsg",
            "type": "u128"
          },
          {
            "name": "linkTokenMint",
            "type": "publicKey"
          },
          {
            "name": "linkTokenLocalDecimals",
            "type": "u8"
          },
          {
            "name": "onramp",
            "type": "publicKey"
          },
          {
            "name": "defaultCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
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
    }
  ],
  "types": [
    {
      "name": "DestChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "isEnabled",
            "type": "bool"
          },
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "maxNumberOfTokensPerMsg",
            "type": "u16"
          },
          {
            "name": "maxDataBytes",
            "type": "u32"
          },
          {
            "name": "maxPerMsgGasLimit",
            "type": "u32"
          },
          {
            "name": "destGasOverhead",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteBase",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteHigh",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteThreshold",
            "type": "u32"
          },
          {
            "name": "destDataAvailabilityOverheadGas",
            "type": "u32"
          },
          {
            "name": "destGasPerDataAvailabilityByte",
            "type": "u16"
          },
          {
            "name": "destDataAvailabilityMultiplierBps",
            "type": "u16"
          },
          {
            "name": "defaultTokenFeeUsdcents",
            "type": "u16"
          },
          {
            "name": "defaultTokenDestGasOverhead",
            "type": "u32"
          },
          {
            "name": "defaultTxGasLimit",
            "type": "u32"
          },
          {
            "name": "gasMultiplierWeiPerEth",
            "type": "u64"
          },
          {
            "name": "networkFeeUsdcents",
            "type": "u32"
          },
          {
            "name": "gasPriceStalenessThreshold",
            "type": "u32"
          },
          {
            "name": "enforceOutOfOrder",
            "type": "bool"
          },
          {
            "name": "chainFamilySelector",
            "type": {
              "array": [
                "u8",
                4
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
            "name": "usdPerUnitGas",
            "type": {
              "defined": "TimestampedPackedU224"
            }
          }
        ]
      }
    },
    {
      "name": "TimestampedPackedU224",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "value",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          },
          {
            "name": "timestamp",
            "type": "i64"
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
      "name": "UsdPerTokenUpdated",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "value",
          "type": {
            "array": [
              "u8",
              28
            ]
          },
          "index": false
        },
        {
          "name": "timestamp",
          "type": "i64",
          "index": false
        }
      ]
    }
  ]
};

export const IDL: FiredrillFeequoter = {
  "version": "0.1.0",
  "name": "firedrill_feequoter",
  "instructions": [
    {
      "name": "initialize",
      "accounts": [
        {
          "name": "feeQuoter",
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
          "isMut": false,
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
          "name": "compound",
          "type": "publicKey"
        }
      ]
    },
    {
      "name": "emitUsdPerTokenUpdated",
      "accounts": [
        {
          "name": "feeQuoter",
          "isMut": true,
          "isSigner": false
        },
        {
          "name": "owner",
          "isMut": false,
          "isSigner": true
        }
      ],
      "args": []
    }
  ],
  "accounts": [
    {
      "name": "firedrillFeeQuoter",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "chainSelector",
            "type": "u64"
          },
          {
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "compound",
            "type": "publicKey"
          },
          {
            "name": "token",
            "type": "publicKey"
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
            "name": "owner",
            "type": "publicKey"
          },
          {
            "name": "proposedOwner",
            "type": "publicKey"
          },
          {
            "name": "maxFeeJuelsPerMsg",
            "type": "u128"
          },
          {
            "name": "linkTokenMint",
            "type": "publicKey"
          },
          {
            "name": "linkTokenLocalDecimals",
            "type": "u8"
          },
          {
            "name": "onramp",
            "type": "publicKey"
          },
          {
            "name": "defaultCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
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
    }
  ],
  "types": [
    {
      "name": "DestChainConfig",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "isEnabled",
            "type": "bool"
          },
          {
            "name": "laneCodeVersion",
            "type": {
              "defined": "CodeVersion"
            }
          },
          {
            "name": "maxNumberOfTokensPerMsg",
            "type": "u16"
          },
          {
            "name": "maxDataBytes",
            "type": "u32"
          },
          {
            "name": "maxPerMsgGasLimit",
            "type": "u32"
          },
          {
            "name": "destGasOverhead",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteBase",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteHigh",
            "type": "u32"
          },
          {
            "name": "destGasPerPayloadByteThreshold",
            "type": "u32"
          },
          {
            "name": "destDataAvailabilityOverheadGas",
            "type": "u32"
          },
          {
            "name": "destGasPerDataAvailabilityByte",
            "type": "u16"
          },
          {
            "name": "destDataAvailabilityMultiplierBps",
            "type": "u16"
          },
          {
            "name": "defaultTokenFeeUsdcents",
            "type": "u16"
          },
          {
            "name": "defaultTokenDestGasOverhead",
            "type": "u32"
          },
          {
            "name": "defaultTxGasLimit",
            "type": "u32"
          },
          {
            "name": "gasMultiplierWeiPerEth",
            "type": "u64"
          },
          {
            "name": "networkFeeUsdcents",
            "type": "u32"
          },
          {
            "name": "gasPriceStalenessThreshold",
            "type": "u32"
          },
          {
            "name": "enforceOutOfOrder",
            "type": "bool"
          },
          {
            "name": "chainFamilySelector",
            "type": {
              "array": [
                "u8",
                4
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
            "name": "usdPerUnitGas",
            "type": {
              "defined": "TimestampedPackedU224"
            }
          }
        ]
      }
    },
    {
      "name": "TimestampedPackedU224",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "value",
            "type": {
              "array": [
                "u8",
                28
              ]
            }
          },
          {
            "name": "timestamp",
            "type": "i64"
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
      "name": "UsdPerTokenUpdated",
      "fields": [
        {
          "name": "token",
          "type": "publicKey",
          "index": false
        },
        {
          "name": "value",
          "type": {
            "array": [
              "u8",
              28
            ]
          },
          "index": false
        },
        {
          "name": "timestamp",
          "type": "i64",
          "index": false
        }
      ]
    }
  ]
};
