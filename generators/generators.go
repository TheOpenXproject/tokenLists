package main

type TGenerationMethods string
type TGeneratorType string

const (
	// GenerationAPI indicates that the list is generated by calling an API
	GenerationAPI TGenerationMethods = "API"
	// GenerationEvents indicates that the list is generated by listening to on-chain events
	GenerationEvents TGenerationMethods = "Events"
	// GenerationExternalList indicates that the list is generated by retrieving a list from an external source
	GenerationExternalList TGenerationMethods = "External"
	// GenerationLegacyList is the same as GenerationExternalList, but it is for deprecated lists
	GenerationLegacyList TGenerationMethods = "Legacy"

	// GeneratorToken indicates that the list is a token list
	GeneratorToken TGeneratorType = "Token"
	// GeneratorPool indicates that the list is a pool list
	GeneratorPool TGeneratorType = "Pool"
)

type TGenerators struct {
	Exec             func()
	Name             string
	Description      string
	GenerationMethod TGenerationMethods
	GeneratorType    TGeneratorType
	Tags             []string //
}

var GENERATORS = map[string]TGenerators{
	`1inch`: {
		Exec:             build1InchTokenList,
		Name:             `1Inch`,
		Description:      `A list of tokens available in 1Inch DeFi / DEX aggregator`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`aerodrome`: {
		Exec:             buildAeroTokenList,
		Name:             `Aerodrome`,
		Description:      `A list of tokens available on Aerodrome.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`bebop`: {
		Exec:             buildBebopTokenList,
		Name:             `Bebop`,
		Description:      `A list of tokens available on Bebop.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`coingecko`: {
		Exec:             buildCoingeckoTokenList,
		Name:             `CoinGecko`,
		Description:      `A list of tokens available showing in CoinGecko data agregator.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`cowswap`: {
		Exec:             buildCowswapTokenList,
		Name:             `Cow Swap`,
		Description:      `A list of tokens available for trading on CoW Swap, a DEX focused on MEV protection.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`curve`: {
		Exec:             buildCurveTokenList,
		Name:             `Curve`,
		Description:      `A list of tokens available for trading on Curve, the largest stableswap.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`defillama`: {
		Exec:             buildDefillamaTokenList,
		Name:             `DefiLlama`,
		Description:      `A list of tokens available in DefiLlama token service`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	// `ethereum-etherscan`: { // deprecated, use scan-1 instead
	// 	Exec:             buildScanTokenList_1,
	// 	Name:             `Ethereum (Etherscan)`,
	// 	Description:      `The top 1000 of tokens available on Ethereum blockchain, retrieved from Etherscan.`,
	// 	GenerationMethod: GenerationAPI,
	// 	GeneratorType:    GeneratorToken,
	// },
	`etherscan`: {
		Exec:             buildScanTokenList,
		Name:             `Etherscan`,
		Description:      `The top of tokens available on by market cap.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`ledger`: {
		Exec:             buildLedgersTokenList,
		Name:             `Ledger`,
		Description:      `A list of tokens supported in Ledger Live App`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`messari`: {
		Exec:             buildMessariTokenList,
		Name:             `Messari`,
		Description:      `A list of tokens registered in Messari`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`optimism`: {
		Exec:             buildOptimismTokenList,
		Name:             `Optimism`,
		Description:      `A list of tokens used as the source of truth for the Optimism Gateway.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`paraswap`: {
		Exec:             buildParaswapTokenList,
		Name:             `Paraswap`,
		Description:      `A list of tokens available for trading on Paraswap DEX`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	// `polygon-zkevm`: { // deprecated, use scan-1101 instead
	// 	Exec:             buildScanTokenList_1101,
	// 	Name:             `Polygon (ZK-EVM)`,
	// 	Description:      `A list of tokens available on the Polygon zkEVM, the first EVM zero-knowledge scaling solution.`,
	// 	GenerationMethod: GenerationAPI,
	// 	GeneratorType:    GeneratorToken,
	// },
	`portals`: {
		Exec:             buildPortalsTokenList,
		Name:             `Portals`,
		Description:      `A list of tokens available for trading on Portals DEX.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	// `scan-8453`: {
	// 	Exec:             buildScanTokenList_8453,
	// 	Name:             `Base (Basescan)`,
	// 	Description:      `The top ERC20 tokens available on Base blockchain, retrieved from Basescan.`,
	// 	GenerationMethod: GenerationAPI,
	// 	GeneratorType:    GeneratorToken,
	// },
	`sushiswap-pairs`: {
		Exec:             buildSushiswapPairsTokenList,
		Name:             `SushiSwap (token pairs)`,
		Description:      `A list of token used in the SushiSwap Liquidity Pools.`,
		GenerationMethod: GenerationEvents,
		GeneratorType:    GeneratorPool,
	},
	`sushiswap-pools`: {
		Exec:             buildSushiswapPoolsTokenList,
		Name:             `SushiSwap (pools)`,
		Description:      `A list of Liquidity Pool available on SushiSwap DEX.`,
		GenerationMethod: GenerationEvents,
		GeneratorType:    GeneratorPool,
	},
	`sushiswap`: {
		Exec:             buildSushiswapTokenList,
		Name:             `SushiSwap`,
		Description:      `A list of tokens available on SushiSwap DEX.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`uniswap-pairs`: {
		Exec:             buildUniswapPairsTokenList,
		Name:             `UniSwap (pairs)`,
		Description:      `A list of token pairs (liquidity pools) available for trading on UniSwap.`,
		GenerationMethod: GenerationEvents,
		GeneratorType:    GeneratorPool,
	},
	`uniswap-pools`: {
		Exec:             buildUniswapPoolsTokenList,
		Name:             `UniSwap (pools)`,
		Description:      `A list of Liquidity Pool available on Uniswap V2 DEX.`,
		GenerationMethod: GenerationEvents,
		GeneratorType:    GeneratorPool,
	},
	`uniswap`: {
		Exec:             buildUniswapTokenList,
		Name:             `UniSwap`,
		Description:      `A list of tokens available on UniSwap DEX.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`velodrome`: {
		Exec:             buildVeloTokenList,
		Name:             `Velodrome`,
		Description:      `A list of tokens available on Velodrome.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`wido`: {
		Exec:             buildWidoTokenList,
		Name:             `Wido`,
		Description:      `A list of tokens supported by the Wido Router`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`yearn-extended`: {
		Exec:             buildYearnExtendedTokenList,
		Name:             `Yearn Extended`,
		Description:      `A list of tokens available for depositing in Yearn, the tokens that represent yVaults and any related tokens.`,
		GenerationMethod: GenerationExternalList,
		GeneratorType:    GeneratorToken,
	},
	`yearn`: {
		Exec:             buildYearnTokenList,
		Name:             `Yearn`,
		Description:      `A list of Yearn's vaults and their underlying tokens.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
	`zksync`: {
		Exec:             buildZkSyncTokenList,
		Name:             `ZKSync`,
		Description:      `A list of zkSync: The future-proof zkEVM.`,
		GenerationMethod: GenerationAPI,
		GeneratorType:    GeneratorToken,
	},
}
