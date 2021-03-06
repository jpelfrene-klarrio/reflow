// Copyright 2017 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// THIS FILE WAS AUTOMATICALLY GENERATED. DO NOT EDIT.

package instances

// Type describes an EC2 instance type.
type Type struct {
	// Name is the API name of this EC2 instance type.
	Name string
	// EBSOptimized is set to true if the instance type permits EBS optimization.
	EBSOptimized bool
	// VCPU stores the number of VCPUs provided by this instance type.
	VCPU uint
	// Memory stores the number of (fractional) GiB of memory provided by this instance type.
	Memory float64
	// Price stores the on-demand price per region for this instance type.
	Price map[string]float64
	// Generation stores the generation name for this instance ("current" or "previous").
	Generation string
}

// Types stores known EC2 instance types.
var Types = []Type{
	{
		Name:         "cc2.8xlarge",
		EBSOptimized: false,
		VCPU:         32,
		Memory:       60.500000,
		Price: map[string]float64{
			"ap-northeast-1": 2.349,
			"eu-west-1":      2.25,
			"us-east-1":      2,
			"us-gov-west-1":  2.25,
			"us-west-2":      2,
		},
		Generation: "previous",
	},
	{
		Name:         "cg1.4xlarge",
		EBSOptimized: false,
		VCPU:         16,
		Memory:       22.500000,
		Price: map[string]float64{
			"eu-west-1": 2.36,
			"us-east-1": 2.1,
		},
		Generation: "previous",
	},
	{
		Name:         "i2.xlarge",
		EBSOptimized: true,
		VCPU:         4,
		Memory:       30.500000,
		Price: map[string]float64{
			"ap-northeast-1": 1.001,
			"ap-northeast-2": 1.001,
			"ap-south-1":     0.967,
			"ap-southeast-1": 1.018,
			"ap-southeast-2": 1.018,
			"eu-central-1":   1.013,
			"eu-west-1":      0.938,
			"us-east-1":      0.853,
			"us-east-2":      0.853,
			"us-gov-west-1":  1.023,
			"us-west-1":      0.938,
			"us-west-2":      0.853,
		},
		Generation: "previous",
	},
	{
		Name:         "i2.2xlarge",
		EBSOptimized: true,
		VCPU:         8,
		Memory:       61.000000,
		Price: map[string]float64{
			"ap-northeast-1": 2.001,
			"ap-northeast-2": 2.001,
			"ap-south-1":     1.933,
			"ap-southeast-1": 2.035,
			"ap-southeast-2": 2.035,
			"eu-central-1":   2.026,
			"eu-west-1":      1.876,
			"us-east-1":      1.705,
			"us-east-2":      1.705,
			"us-gov-west-1":  2.046,
			"us-west-1":      1.876,
			"us-west-2":      1.705,
		},
		Generation: "previous",
	},
	{
		Name:         "i2.4xlarge",
		EBSOptimized: true,
		VCPU:         16,
		Memory:       122.000000,
		Price: map[string]float64{
			"ap-northeast-1": 4.002,
			"ap-northeast-2": 4.002,
			"ap-south-1":     3.867,
			"ap-southeast-1": 4.07,
			"ap-southeast-2": 4.07,
			"eu-central-1":   4.051,
			"eu-west-1":      3.751,
			"us-east-1":      3.41,
			"us-east-2":      3.41,
			"us-gov-west-1":  4.092,
			"us-west-1":      3.751,
			"us-west-2":      3.41,
		},
		Generation: "previous",
	},
	{
		Name:         "i2.8xlarge",
		EBSOptimized: true,
		VCPU:         32,
		Memory:       244.000000,
		Price: map[string]float64{
			"ap-northeast-1": 8.004,
			"ap-northeast-2": 8.004,
			"ap-south-1":     7.733,
			"ap-southeast-1": 8.14,
			"ap-southeast-2": 8.14,
			"eu-central-1":   8.102,
			"eu-west-1":      7.502,
			"us-east-1":      6.82,
			"us-east-2":      6.82,
			"us-gov-west-1":  8.184,
			"us-west-1":      7.502,
			"us-west-2":      6.82,
		},
		Generation: "previous",
	},
	{
		Name:         "hi1.4xlarge",
		EBSOptimized: false,
		VCPU:         16,
		Memory:       60.500000,
		Price: map[string]float64{
			"ap-northeast-1": 3.276,
			"eu-west-1":      3.1,
			"us-east-1":      3.1,
			"us-west-2":      3.1,
		},
		Generation: "previous",
	},
	{
		Name:         "hs1.8xlarge",
		EBSOptimized: false,
		VCPU:         16,
		Memory:       117.000000,
		Price: map[string]float64{
			"ap-northeast-1": 5.4,
			"ap-southeast-1": 5.57,
			"ap-southeast-2": 5.57,
			"eu-west-1":      4.9,
			"us-east-1":      4.6,
			"us-gov-west-1":  5.52,
			"us-west-2":      4.6,
		},
		Generation: "previous",
	},
	{
		Name:         "t2.nano",
		EBSOptimized: false,
		VCPU:         1,
		Memory:       0.500000,
		Price: map[string]float64{
			"ap-northeast-1": 0.008,
			"ap-northeast-2": 0.008,
			"ap-south-1":     0.0074,
			"ap-southeast-1": 0.0075,
			"ap-southeast-2": 0.008,
			"ca-central-1":   0.0065,
			"eu-central-1":   0.0068,
			"eu-west-1":      0.0063,
			"eu-west-2":      0.0066,
			"sa-east-1":      0.0101,
			"us-east-1":      0.0059,
			"us-east-2":      0.0059,
			"us-gov-west-1":  0.0068,
			"us-west-1":      0.0077,
			"us-west-2":      0.0059,
		},
		Generation: "current",
	},
	{
		Name:         "t2.micro",
		EBSOptimized: false,
		VCPU:         1,
		Memory:       1.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.016,
			"ap-northeast-2": 0.016,
			"ap-south-1":     0.015,
			"ap-southeast-1": 0.015,
			"ap-southeast-2": 0.016,
			"ca-central-1":   0.013,
			"eu-central-1":   0.014,
			"eu-west-1":      0.013,
			"eu-west-2":      0.014,
			"sa-east-1":      0.02,
			"us-east-1":      0.012,
			"us-east-2":      0.012,
			"us-gov-west-1":  0.014,
			"us-west-1":      0.015,
			"us-west-2":      0.012,
		},
		Generation: "current",
	},
	{
		Name:         "t2.small",
		EBSOptimized: false,
		VCPU:         1,
		Memory:       2.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.032,
			"ap-northeast-2": 0.032,
			"ap-south-1":     0.03,
			"ap-southeast-1": 0.03,
			"ap-southeast-2": 0.032,
			"ca-central-1":   0.026,
			"eu-central-1":   0.027,
			"eu-west-1":      0.025,
			"eu-west-2":      0.026,
			"sa-east-1":      0.041,
			"us-east-1":      0.023,
			"us-east-2":      0.023,
			"us-gov-west-1":  0.028,
			"us-west-1":      0.031,
			"us-west-2":      0.023,
		},
		Generation: "current",
	},
	{
		Name:         "t2.medium",
		EBSOptimized: false,
		VCPU:         2,
		Memory:       4.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.064,
			"ap-northeast-2": 0.064,
			"ap-south-1":     0.059,
			"ap-southeast-1": 0.06,
			"ap-southeast-2": 0.064,
			"ca-central-1":   0.052,
			"eu-central-1":   0.054,
			"eu-west-1":      0.05,
			"eu-west-2":      0.052,
			"sa-east-1":      0.081,
			"us-east-1":      0.047,
			"us-east-2":      0.047,
			"us-gov-west-1":  0.056,
			"us-west-1":      0.061,
			"us-west-2":      0.047,
		},
		Generation: "current",
	},
	{
		Name:         "t2.large",
		EBSOptimized: false,
		VCPU:         2,
		Memory:       8.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.128,
			"ap-northeast-2": 0.128,
			"ap-south-1":     0.119,
			"ap-southeast-1": 0.12,
			"ap-southeast-2": 0.128,
			"ca-central-1":   0.103,
			"eu-central-1":   0.108,
			"eu-west-1":      0.101,
			"eu-west-2":      0.106,
			"sa-east-1":      0.162,
			"us-east-1":      0.094,
			"us-east-2":      0.094,
			"us-gov-west-1":  0.112,
			"us-west-1":      0.122,
			"us-west-2":      0.094,
		},
		Generation: "current",
	},
	{
		Name:         "t2.xlarge",
		EBSOptimized: false,
		VCPU:         4,
		Memory:       16.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.256,
			"ap-northeast-2": 0.256,
			"ap-south-1":     0.238,
			"ap-southeast-1": 0.24,
			"ap-southeast-2": 0.256,
			"ca-central-1":   0.206,
			"eu-central-1":   0.216,
			"eu-west-1":      0.202,
			"eu-west-2":      0.212,
			"sa-east-1":      0.324,
			"us-east-1":      0.188,
			"us-east-2":      0.188,
			"us-gov-west-1":  0.224,
			"us-west-1":      0.244,
			"us-west-2":      0.188,
		},
		Generation: "current",
	},
	{
		Name:         "t2.2xlarge",
		EBSOptimized: false,
		VCPU:         8,
		Memory:       32.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.512,
			"ap-northeast-2": 0.512,
			"ap-south-1":     0.476,
			"ap-southeast-1": 0.48,
			"ap-southeast-2": 0.512,
			"ca-central-1":   0.412,
			"eu-central-1":   0.432,
			"eu-west-1":      0.404,
			"eu-west-2":      0.424,
			"sa-east-1":      0.648,
			"us-east-1":      0.376,
			"us-east-2":      0.376,
			"us-gov-west-1":  0.448,
			"us-west-1":      0.488,
			"us-west-2":      0.376,
		},
		Generation: "current",
	},
	{
		Name:         "m4.large",
		EBSOptimized: true,
		VCPU:         2,
		Memory:       8.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.129,
			"ap-northeast-2": 0.123,
			"ap-south-1":     0.123,
			"ap-southeast-1": 0.125,
			"ap-southeast-2": 0.125,
			"ca-central-1":   0.111,
			"eu-central-1":   0.12,
			"eu-west-1":      0.111,
			"eu-west-2":      0.116,
			"sa-east-1":      0.159,
			"us-east-1":      0.1,
			"us-east-2":      0.1,
			"us-gov-west-1":  0.126,
			"us-west-1":      0.117,
			"us-west-2":      0.1,
		},
		Generation: "current",
	},
	{
		Name:         "m4.xlarge",
		EBSOptimized: true,
		VCPU:         4,
		Memory:       16.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.258,
			"ap-northeast-2": 0.246,
			"ap-south-1":     0.246,
			"ap-southeast-1": 0.25,
			"ap-southeast-2": 0.25,
			"ca-central-1":   0.222,
			"eu-central-1":   0.24,
			"eu-west-1":      0.222,
			"eu-west-2":      0.232,
			"sa-east-1":      0.318,
			"us-east-1":      0.2,
			"us-east-2":      0.2,
			"us-gov-west-1":  0.252,
			"us-west-1":      0.234,
			"us-west-2":      0.2,
		},
		Generation: "current",
	},
	{
		Name:         "m4.2xlarge",
		EBSOptimized: true,
		VCPU:         8,
		Memory:       32.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.516,
			"ap-northeast-2": 0.492,
			"ap-south-1":     0.492,
			"ap-southeast-1": 0.5,
			"ap-southeast-2": 0.5,
			"ca-central-1":   0.444,
			"eu-central-1":   0.48,
			"eu-west-1":      0.444,
			"eu-west-2":      0.464,
			"sa-east-1":      0.636,
			"us-east-1":      0.4,
			"us-east-2":      0.4,
			"us-gov-west-1":  0.504,
			"us-west-1":      0.468,
			"us-west-2":      0.4,
		},
		Generation: "current",
	},
	{
		Name:         "m4.4xlarge",
		EBSOptimized: true,
		VCPU:         16,
		Memory:       64.000000,
		Price: map[string]float64{
			"ap-northeast-1": 1.032,
			"ap-northeast-2": 0.984,
			"ap-south-1":     0.984,
			"ap-southeast-1": 1,
			"ap-southeast-2": 1,
			"ca-central-1":   0.888,
			"eu-central-1":   0.96,
			"eu-west-1":      0.888,
			"eu-west-2":      0.928,
			"sa-east-1":      1.272,
			"us-east-1":      0.8,
			"us-east-2":      0.8,
			"us-gov-west-1":  1.008,
			"us-west-1":      0.936,
			"us-west-2":      0.8,
		},
		Generation: "current",
	},
	{
		Name:         "m4.10xlarge",
		EBSOptimized: true,
		VCPU:         40,
		Memory:       160.000000,
		Price: map[string]float64{
			"ap-northeast-1": 2.58,
			"ap-northeast-2": 2.46,
			"ap-south-1":     2.46,
			"ap-southeast-1": 2.5,
			"ap-southeast-2": 2.5,
			"ca-central-1":   2.22,
			"eu-central-1":   2.4,
			"eu-west-1":      2.22,
			"eu-west-2":      2.32,
			"sa-east-1":      3.18,
			"us-east-1":      2,
			"us-east-2":      2,
			"us-gov-west-1":  2.52,
			"us-west-1":      2.34,
			"us-west-2":      2,
		},
		Generation: "current",
	},
	{
		Name:         "m4.16xlarge",
		EBSOptimized: true,
		VCPU:         64,
		Memory:       256.000000,
		Price: map[string]float64{
			"ap-northeast-1": 4.128,
			"ap-northeast-2": 3.936,
			"ap-south-1":     3.936,
			"ap-southeast-1": 4,
			"ap-southeast-2": 4,
			"ca-central-1":   3.552,
			"eu-central-1":   3.84,
			"eu-west-1":      3.552,
			"eu-west-2":      3.712,
			"sa-east-1":      5.088,
			"us-east-1":      3.2,
			"us-east-2":      3.2,
			"us-gov-west-1":  4.032,
			"us-west-1":      3.744,
			"us-west-2":      3.2,
		},
		Generation: "current",
	},
	{
		Name:         "m3.medium",
		EBSOptimized: false,
		VCPU:         1,
		Memory:       3.750000,
		Price: map[string]float64{
			"ap-northeast-1": 0.096,
			"ap-southeast-1": 0.098,
			"ap-southeast-2": 0.093,
			"eu-central-1":   0.079,
			"eu-west-1":      0.073,
			"sa-east-1":      0.095,
			"us-east-1":      0.067,
			"us-gov-west-1":  0.084,
			"us-west-1":      0.077,
			"us-west-2":      0.067,
		},
		Generation: "current",
	},
	{
		Name:         "m3.large",
		EBSOptimized: false,
		VCPU:         2,
		Memory:       7.500000,
		Price: map[string]float64{
			"ap-northeast-1": 0.193,
			"ap-southeast-1": 0.196,
			"ap-southeast-2": 0.186,
			"eu-central-1":   0.158,
			"eu-west-1":      0.146,
			"sa-east-1":      0.19,
			"us-east-1":      0.133,
			"us-gov-west-1":  0.168,
			"us-west-1":      0.154,
			"us-west-2":      0.133,
		},
		Generation: "current",
	},
	{
		Name:         "m3.xlarge",
		EBSOptimized: true,
		VCPU:         4,
		Memory:       15.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.385,
			"ap-southeast-1": 0.392,
			"ap-southeast-2": 0.372,
			"eu-central-1":   0.315,
			"eu-west-1":      0.293,
			"sa-east-1":      0.381,
			"us-east-1":      0.266,
			"us-gov-west-1":  0.336,
			"us-west-1":      0.308,
			"us-west-2":      0.266,
		},
		Generation: "current",
	},
	{
		Name:         "m3.2xlarge",
		EBSOptimized: true,
		VCPU:         8,
		Memory:       30.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.77,
			"ap-southeast-1": 0.784,
			"ap-southeast-2": 0.745,
			"eu-central-1":   0.632,
			"eu-west-1":      0.585,
			"sa-east-1":      0.761,
			"us-east-1":      0.532,
			"us-gov-west-1":  0.672,
			"us-west-1":      0.616,
			"us-west-2":      0.532,
		},
		Generation: "current",
	},
	{
		Name:         "c4.large",
		EBSOptimized: true,
		VCPU:         2,
		Memory:       3.750000,
		Price: map[string]float64{
			"ap-northeast-1": 0.126,
			"ap-northeast-2": 0.114,
			"ap-south-1":     0.11,
			"ap-southeast-1": 0.115,
			"ap-southeast-2": 0.13,
			"ca-central-1":   0.11,
			"eu-central-1":   0.114,
			"eu-west-1":      0.113,
			"eu-west-2":      0.119,
			"sa-east-1":      0.155,
			"us-east-1":      0.1,
			"us-east-2":      0.1,
			"us-gov-west-1":  0.12,
			"us-west-1":      0.124,
			"us-west-2":      0.1,
		},
		Generation: "current",
	},
	{
		Name:         "c4.xlarge",
		EBSOptimized: true,
		VCPU:         4,
		Memory:       7.500000,
		Price: map[string]float64{
			"ap-northeast-1": 0.252,
			"ap-northeast-2": 0.227,
			"ap-south-1":     0.22,
			"ap-southeast-1": 0.231,
			"ap-southeast-2": 0.261,
			"ca-central-1":   0.218,
			"eu-central-1":   0.227,
			"eu-west-1":      0.226,
			"eu-west-2":      0.237,
			"sa-east-1":      0.309,
			"us-east-1":      0.199,
			"us-east-2":      0.199,
			"us-gov-west-1":  0.239,
			"us-west-1":      0.249,
			"us-west-2":      0.199,
		},
		Generation: "current",
	},
	{
		Name:         "c4.2xlarge",
		EBSOptimized: true,
		VCPU:         8,
		Memory:       15.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.504,
			"ap-northeast-2": 0.454,
			"ap-south-1":     0.439,
			"ap-southeast-1": 0.462,
			"ap-southeast-2": 0.522,
			"ca-central-1":   0.438,
			"eu-central-1":   0.454,
			"eu-west-1":      0.453,
			"eu-west-2":      0.476,
			"sa-east-1":      0.618,
			"us-east-1":      0.398,
			"us-east-2":      0.398,
			"us-gov-west-1":  0.479,
			"us-west-1":      0.498,
			"us-west-2":      0.398,
		},
		Generation: "current",
	},
	{
		Name:         "c4.4xlarge",
		EBSOptimized: true,
		VCPU:         16,
		Memory:       30.000000,
		Price: map[string]float64{
			"ap-northeast-1": 1.008,
			"ap-northeast-2": 0.907,
			"ap-south-1":     0.878,
			"ap-southeast-1": 0.924,
			"ap-southeast-2": 1.042,
			"ca-central-1":   0.876,
			"eu-central-1":   0.909,
			"eu-west-1":      0.905,
			"eu-west-2":      0.95,
			"sa-east-1":      1.235,
			"us-east-1":      0.796,
			"us-east-2":      0.796,
			"us-gov-west-1":  0.958,
			"us-west-1":      0.997,
			"us-west-2":      0.796,
		},
		Generation: "current",
	},
	{
		Name:         "c4.8xlarge",
		EBSOptimized: true,
		VCPU:         36,
		Memory:       60.000000,
		Price: map[string]float64{
			"ap-northeast-1": 2.016,
			"ap-northeast-2": 1.815,
			"ap-south-1":     1.756,
			"ap-southeast-1": 1.848,
			"ap-southeast-2": 2.085,
			"ca-central-1":   1.75,
			"eu-central-1":   1.817,
			"eu-west-1":      1.811,
			"eu-west-2":      1.902,
			"sa-east-1":      2.47,
			"us-east-1":      1.591,
			"us-east-2":      1.591,
			"us-gov-west-1":  1.915,
			"us-west-1":      1.993,
			"us-west-2":      1.591,
		},
		Generation: "current",
	},
	{
		Name:         "c3.large",
		EBSOptimized: false,
		VCPU:         2,
		Memory:       3.750000,
		Price: map[string]float64{
			"ap-northeast-1": 0.128,
			"ap-southeast-1": 0.132,
			"ap-southeast-2": 0.132,
			"eu-central-1":   0.129,
			"eu-west-1":      0.12,
			"sa-east-1":      0.163,
			"us-east-1":      0.105,
			"us-gov-west-1":  0.126,
			"us-west-1":      0.12,
			"us-west-2":      0.105,
		},
		Generation: "current",
	},
	{
		Name:         "c3.xlarge",
		EBSOptimized: true,
		VCPU:         4,
		Memory:       7.500000,
		Price: map[string]float64{
			"ap-northeast-1": 0.255,
			"ap-southeast-1": 0.265,
			"ap-southeast-2": 0.265,
			"eu-central-1":   0.258,
			"eu-west-1":      0.239,
			"sa-east-1":      0.325,
			"us-east-1":      0.21,
			"us-gov-west-1":  0.252,
			"us-west-1":      0.239,
			"us-west-2":      0.21,
		},
		Generation: "current",
	},
	{
		Name:         "c3.2xlarge",
		EBSOptimized: true,
		VCPU:         8,
		Memory:       15.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.511,
			"ap-southeast-1": 0.529,
			"ap-southeast-2": 0.529,
			"eu-central-1":   0.516,
			"eu-west-1":      0.478,
			"sa-east-1":      0.65,
			"us-east-1":      0.42,
			"us-gov-west-1":  0.504,
			"us-west-1":      0.478,
			"us-west-2":      0.42,
		},
		Generation: "current",
	},
	{
		Name:         "c3.4xlarge",
		EBSOptimized: true,
		VCPU:         16,
		Memory:       30.000000,
		Price: map[string]float64{
			"ap-northeast-1": 1.021,
			"ap-southeast-1": 1.058,
			"ap-southeast-2": 1.058,
			"eu-central-1":   1.032,
			"eu-west-1":      0.956,
			"sa-east-1":      1.3,
			"us-east-1":      0.84,
			"us-gov-west-1":  1.008,
			"us-west-1":      0.956,
			"us-west-2":      0.84,
		},
		Generation: "current",
	},
	{
		Name:         "c3.8xlarge",
		EBSOptimized: false,
		VCPU:         32,
		Memory:       60.000000,
		Price: map[string]float64{
			"ap-northeast-1": 2.043,
			"ap-southeast-1": 2.117,
			"ap-southeast-2": 2.117,
			"eu-central-1":   2.064,
			"eu-west-1":      1.912,
			"sa-east-1":      2.6,
			"us-east-1":      1.68,
			"us-gov-west-1":  2.016,
			"us-west-1":      1.912,
			"us-west-2":      1.68,
		},
		Generation: "current",
	},
	{
		Name:         "x1.16xlarge",
		EBSOptimized: true,
		VCPU:         64,
		Memory:       976.000000,
		Price: map[string]float64{
			"ap-northeast-1": 9.671,
			"ap-northeast-2": 9.671,
			"ap-south-1":     9.187,
			"ap-southeast-1": 9.671,
			"ap-southeast-2": 9.671,
			"ca-central-1":   7.336,
			"eu-central-1":   9.337,
			"eu-west-1":      8.003,
			"eu-west-2":      8.403,
			"sa-east-1":      13.005,
			"us-east-1":      6.669,
			"us-east-2":      6.669,
			"us-gov-west-1":  8.003,
			"us-west-2":      6.669,
		},
		Generation: "current",
	},
	{
		Name:         "x1.32xlarge",
		EBSOptimized: true,
		VCPU:         128,
		Memory:       1952.000000,
		Price: map[string]float64{
			"ap-northeast-1": 19.341,
			"ap-northeast-2": 19.341,
			"ap-south-1":     18.374,
			"ap-southeast-1": 19.341,
			"ap-southeast-2": 19.341,
			"ca-central-1":   14.672,
			"eu-central-1":   18.674,
			"eu-west-1":      16.006,
			"eu-west-2":      16.806,
			"sa-east-1":      26.01,
			"us-east-1":      13.338,
			"us-east-2":      13.338,
			"us-gov-west-1":  16.006,
			"us-west-2":      13.338,
		},
		Generation: "current",
	},
	{
		Name:         "r4.large",
		EBSOptimized: true,
		VCPU:         2,
		Memory:       15.250000,
		Price: map[string]float64{
			"ap-northeast-1": 0.16,
			"ap-northeast-2": 0.16,
			"ap-south-1":     0.152,
			"ap-southeast-1": 0.16,
			"ap-southeast-2": 0.16,
			"ca-central-1":   0.146,
			"eu-central-1":   0.16,
			"eu-west-1":      0.148,
			"eu-west-2":      0.156,
			"sa-east-1":      0.28,
			"us-east-1":      0.133,
			"us-east-2":      0.133,
			"us-gov-west-1":  0.16,
			"us-west-1":      0.148,
			"us-west-2":      0.133,
		},
		Generation: "current",
	},
	{
		Name:         "r4.xlarge",
		EBSOptimized: true,
		VCPU:         4,
		Memory:       30.500000,
		Price: map[string]float64{
			"ap-northeast-1": 0.32,
			"ap-northeast-2": 0.32,
			"ap-south-1":     0.304,
			"ap-southeast-1": 0.32,
			"ap-southeast-2": 0.319,
			"ca-central-1":   0.292,
			"eu-central-1":   0.32,
			"eu-west-1":      0.296,
			"eu-west-2":      0.312,
			"sa-east-1":      0.56,
			"us-east-1":      0.266,
			"us-east-2":      0.266,
			"us-gov-west-1":  0.319,
			"us-west-1":      0.296,
			"us-west-2":      0.266,
		},
		Generation: "current",
	},
	{
		Name:         "r4.2xlarge",
		EBSOptimized: true,
		VCPU:         8,
		Memory:       61.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.64,
			"ap-northeast-2": 0.64,
			"ap-south-1":     0.608,
			"ap-southeast-1": 0.64,
			"ap-southeast-2": 0.638,
			"ca-central-1":   0.584,
			"eu-central-1":   0.64,
			"eu-west-1":      0.593,
			"eu-west-2":      0.624,
			"sa-east-1":      1.12,
			"us-east-1":      0.532,
			"us-east-2":      0.532,
			"us-gov-west-1":  0.638,
			"us-west-1":      0.593,
			"us-west-2":      0.532,
		},
		Generation: "current",
	},
	{
		Name:         "r4.4xlarge",
		EBSOptimized: true,
		VCPU:         16,
		Memory:       122.000000,
		Price: map[string]float64{
			"ap-northeast-1": 1.28,
			"ap-northeast-2": 1.28,
			"ap-south-1":     1.216,
			"ap-southeast-1": 1.28,
			"ap-southeast-2": 1.277,
			"ca-central-1":   1.168,
			"eu-central-1":   1.28,
			"eu-west-1":      1.186,
			"eu-west-2":      1.248,
			"sa-east-1":      2.24,
			"us-east-1":      1.064,
			"us-east-2":      1.064,
			"us-gov-west-1":  1.277,
			"us-west-1":      1.186,
			"us-west-2":      1.064,
		},
		Generation: "current",
	},
	{
		Name:         "r4.8xlarge",
		EBSOptimized: true,
		VCPU:         32,
		Memory:       244.000000,
		Price: map[string]float64{
			"ap-northeast-1": 2.56,
			"ap-northeast-2": 2.56,
			"ap-south-1":     2.432,
			"ap-southeast-1": 2.56,
			"ap-southeast-2": 2.554,
			"ca-central-1":   2.336,
			"eu-central-1":   2.561,
			"eu-west-1":      2.371,
			"eu-west-2":      2.496,
			"sa-east-1":      4.48,
			"us-east-1":      2.128,
			"us-east-2":      2.128,
			"us-gov-west-1":  2.554,
			"us-west-1":      2.371,
			"us-west-2":      2.128,
		},
		Generation: "current",
	},
	{
		Name:         "r4.16xlarge",
		EBSOptimized: true,
		VCPU:         64,
		Memory:       488.000000,
		Price: map[string]float64{
			"ap-northeast-1": 5.12,
			"ap-northeast-2": 5.12,
			"ap-south-1":     4.864,
			"ap-southeast-1": 5.12,
			"ap-southeast-2": 5.107,
			"ca-central-1":   4.672,
			"eu-central-1":   5.122,
			"eu-west-1":      4.742,
			"eu-west-2":      4.992,
			"sa-east-1":      8.96,
			"us-east-1":      4.256,
			"us-east-2":      4.256,
			"us-gov-west-1":  5.107,
			"us-west-1":      4.742,
			"us-west-2":      4.256,
		},
		Generation: "current",
	},
	{
		Name:         "r3.large",
		EBSOptimized: false,
		VCPU:         2,
		Memory:       15.250000,
		Price: map[string]float64{
			"ap-northeast-1": 0.2,
			"ap-northeast-2": 0.2,
			"ap-south-1":     0.19,
			"ap-southeast-1": 0.2,
			"ap-southeast-2": 0.2,
			"eu-central-1":   0.2,
			"eu-west-1":      0.185,
			"sa-east-1":      0.35,
			"us-east-1":      0.166,
			"us-east-2":      0.166,
			"us-gov-west-1":  0.2,
			"us-west-1":      0.185,
			"us-west-2":      0.166,
		},
		Generation: "current",
	},
	{
		Name:         "r3.xlarge",
		EBSOptimized: true,
		VCPU:         4,
		Memory:       30.500000,
		Price: map[string]float64{
			"ap-northeast-1": 0.399,
			"ap-northeast-2": 0.399,
			"ap-south-1":     0.379,
			"ap-southeast-1": 0.399,
			"ap-southeast-2": 0.399,
			"eu-central-1":   0.4,
			"eu-west-1":      0.371,
			"sa-east-1":      0.7,
			"us-east-1":      0.333,
			"us-east-2":      0.333,
			"us-gov-west-1":  0.399,
			"us-west-1":      0.371,
			"us-west-2":      0.333,
		},
		Generation: "current",
	},
	{
		Name:         "r3.2xlarge",
		EBSOptimized: true,
		VCPU:         8,
		Memory:       61.000000,
		Price: map[string]float64{
			"ap-northeast-1": 0.798,
			"ap-northeast-2": 0.798,
			"ap-south-1":     0.758,
			"ap-southeast-1": 0.798,
			"ap-southeast-2": 0.798,
			"eu-central-1":   0.8,
			"eu-west-1":      0.741,
			"sa-east-1":      1.399,
			"us-east-1":      0.665,
			"us-east-2":      0.665,
			"us-gov-west-1":  0.798,
			"us-west-1":      0.741,
			"us-west-2":      0.665,
		},
		Generation: "current",
	},
	{
		Name:         "r3.4xlarge",
		EBSOptimized: true,
		VCPU:         16,
		Memory:       122.000000,
		Price: map[string]float64{
			"ap-northeast-1": 1.596,
			"ap-northeast-2": 1.596,
			"ap-south-1":     1.516,
			"ap-southeast-1": 1.596,
			"ap-southeast-2": 1.596,
			"eu-central-1":   1.6,
			"eu-west-1":      1.482,
			"sa-east-1":      2.799,
			"us-east-1":      1.33,
			"us-east-2":      1.33,
			"us-gov-west-1":  1.596,
			"us-west-1":      1.482,
			"us-west-2":      1.33,
		},
		Generation: "current",
	},
	{
		Name:         "r3.8xlarge",
		EBSOptimized: false,
		VCPU:         32,
		Memory:       244.000000,
		Price: map[string]float64{
			"ap-northeast-1": 3.192,
			"ap-northeast-2": 3.192,
			"ap-south-1":     3.032,
			"ap-southeast-1": 3.192,
			"ap-southeast-2": 3.192,
			"eu-central-1":   3.201,
			"eu-west-1":      2.964,
			"sa-east-1":      5.597,
			"us-east-1":      2.66,
			"us-east-2":      2.66,
			"us-gov-west-1":  3.192,
			"us-west-1":      2.964,
			"us-west-2":      2.66,
		},
		Generation: "current",
	},
	{
		Name:         "p2.xlarge",
		EBSOptimized: true,
		VCPU:         4,
		Memory:       61.000000,
		Price: map[string]float64{
			"ap-northeast-1": 1.542,
			"ap-northeast-2": 1.465,
			"ap-south-1":     1.718,
			"ap-southeast-1": 1.718,
			"ap-southeast-2": 1.542,
			"eu-central-1":   1.326,
			"eu-west-1":      0.972,
			"us-east-1":      0.9,
			"us-east-2":      0.9,
			"us-gov-west-1":  1.08,
			"us-west-2":      0.9,
		},
		Generation: "current",
	},
	{
		Name:         "p2.8xlarge",
		EBSOptimized: true,
		VCPU:         32,
		Memory:       488.000000,
		Price: map[string]float64{
			"ap-northeast-1": 12.336,
			"ap-northeast-2": 11.72,
			"ap-south-1":     13.744,
			"ap-southeast-1": 13.744,
			"ap-southeast-2": 12.336,
			"eu-central-1":   10.608,
			"eu-west-1":      7.776,
			"us-east-1":      7.2,
			"us-east-2":      7.2,
			"us-gov-west-1":  8.64,
			"us-west-2":      7.2,
		},
		Generation: "current",
	},
	{
		Name:         "p2.16xlarge",
		EBSOptimized: true,
		VCPU:         64,
		Memory:       732.000000,
		Price: map[string]float64{
			"ap-northeast-1": 24.672,
			"ap-northeast-2": 23.44,
			"ap-south-1":     27.488,
			"ap-southeast-1": 27.488,
			"ap-southeast-2": 24.672,
			"eu-central-1":   21.216,
			"eu-west-1":      15.552,
			"us-east-1":      14.4,
			"us-east-2":      14.4,
			"us-gov-west-1":  17.28,
			"us-west-2":      14.4,
		},
		Generation: "current",
	},
	{
		Name:         "g3.4xlarge",
		EBSOptimized: true,
		VCPU:         16,
		Memory:       122.000000,
		Price: map[string]float64{
			"ap-northeast-1": 1.58,
			"ap-southeast-1": 1.67,
			"ap-southeast-2": 1.754,
			"eu-central-1":   1.425,
			"eu-west-1":      1.21,
			"us-east-1":      1.14,
			"us-east-2":      1.14,
			"us-gov-west-1":  1.32,
			"us-west-1":      1.534,
			"us-west-2":      1.14,
		},
		Generation: "current",
	},
	{
		Name:         "g3.8xlarge",
		EBSOptimized: true,
		VCPU:         32,
		Memory:       244.000000,
		Price: map[string]float64{
			"ap-northeast-1": 3.16,
			"ap-southeast-1": 3.34,
			"ap-southeast-2": 3.508,
			"eu-central-1":   2.85,
			"eu-west-1":      2.42,
			"us-east-1":      2.28,
			"us-east-2":      2.28,
			"us-gov-west-1":  2.64,
			"us-west-1":      3.068,
			"us-west-2":      2.28,
		},
		Generation: "current",
	},
	{
		Name:         "g3.16xlarge",
		EBSOptimized: true,
		VCPU:         64,
		Memory:       488.000000,
		Price: map[string]float64{
			"ap-northeast-1": 6.32,
			"ap-southeast-1": 6.68,
			"ap-southeast-2": 7.016,
			"eu-central-1":   5.7,
			"eu-west-1":      4.84,
			"us-east-1":      4.56,
			"us-east-2":      4.56,
			"us-gov-west-1":  5.28,
			"us-west-1":      6.136,
			"us-west-2":      4.56,
		},
		Generation: "current",
	},
	{
		Name:         "f1.2xlarge",
		EBSOptimized: true,
		VCPU:         8,
		Memory:       122.000000,
		Price: map[string]float64{
			"eu-west-1": 1.815,
			"us-east-1": 1.65,
			"us-west-2": 1.65,
		},
		Generation: "current",
	},
	{
		Name:         "f1.16xlarge",
		EBSOptimized: true,
		VCPU:         64,
		Memory:       976.000000,
		Price: map[string]float64{
			"eu-west-1": 14.52,
			"us-east-1": 13.2,
			"us-west-2": 13.2,
		},
		Generation: "current",
	},
	{
		Name:         "d2.xlarge",
		EBSOptimized: true,
		VCPU:         4,
		Memory:       30.500000,
		Price: map[string]float64{
			"ap-northeast-1": 0.844,
			"ap-northeast-2": 0.844,
			"ap-south-1":     0.827,
			"ap-southeast-1": 0.87,
			"ap-southeast-2": 0.87,
			"ca-central-1":   0.759,
			"eu-central-1":   0.794,
			"eu-west-1":      0.735,
			"eu-west-2":      0.772,
			"us-east-1":      0.69,
			"us-east-2":      0.69,
			"us-gov-west-1":  0.828,
			"us-west-1":      0.781,
			"us-west-2":      0.69,
		},
		Generation: "current",
	},
	{
		Name:         "d2.2xlarge",
		EBSOptimized: true,
		VCPU:         8,
		Memory:       61.000000,
		Price: map[string]float64{
			"ap-northeast-1": 1.688,
			"ap-northeast-2": 1.688,
			"ap-south-1":     1.653,
			"ap-southeast-1": 1.74,
			"ap-southeast-2": 1.74,
			"ca-central-1":   1.518,
			"eu-central-1":   1.588,
			"eu-west-1":      1.47,
			"eu-west-2":      1.544,
			"us-east-1":      1.38,
			"us-east-2":      1.38,
			"us-gov-west-1":  1.656,
			"us-west-1":      1.563,
			"us-west-2":      1.38,
		},
		Generation: "current",
	},
	{
		Name:         "d2.4xlarge",
		EBSOptimized: true,
		VCPU:         16,
		Memory:       122.000000,
		Price: map[string]float64{
			"ap-northeast-1": 3.376,
			"ap-northeast-2": 3.376,
			"ap-south-1":     3.306,
			"ap-southeast-1": 3.48,
			"ap-southeast-2": 3.48,
			"ca-central-1":   3.036,
			"eu-central-1":   3.176,
			"eu-west-1":      2.94,
			"eu-west-2":      3.087,
			"us-east-1":      2.76,
			"us-east-2":      2.76,
			"us-gov-west-1":  3.312,
			"us-west-1":      3.125,
			"us-west-2":      2.76,
		},
		Generation: "current",
	},
	{
		Name:         "d2.8xlarge",
		EBSOptimized: true,
		VCPU:         36,
		Memory:       244.000000,
		Price: map[string]float64{
			"ap-northeast-1": 6.752,
			"ap-northeast-2": 6.752,
			"ap-south-1":     6.612,
			"ap-southeast-1": 6.96,
			"ap-southeast-2": 6.96,
			"ca-central-1":   6.072,
			"eu-central-1":   6.352,
			"eu-west-1":      5.88,
			"eu-west-2":      6.174,
			"us-east-1":      5.52,
			"us-east-2":      5.52,
			"us-gov-west-1":  6.624,
			"us-west-1":      6.25,
			"us-west-2":      5.52,
		},
		Generation: "current",
	},
}
