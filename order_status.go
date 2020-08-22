package binance

// OrderStatus represents order status enum.
type OrderStatus string

// OrderType represents order type enum.
type OrderType string

// OrderSide represents order side enum.
type OrderSide string

// OrderStatus, OrderType and OrderSide variables
// make sure to get the terms correct.
var (
	StatusNew             = OrderStatus("NEW")
	StatusPartiallyFilled = OrderStatus("PARTIALLY_FILLED")
	StatusFilled          = OrderStatus("FILLED")
	StatusCancelled       = OrderStatus("CANCELED")
	StatusPendingCancel   = OrderStatus("PENDING_CANCEL")
	StatusRejected        = OrderStatus("REJECTED")
	StatusExpired         = OrderStatus("EXPIRED")

	TypeLimit           = OrderType("LIMIT")
	TypeMarket          = OrderType("MARKET")
	TypeStopLoss        = OrderType("STOP_LOSS")
	TypeStopLossLimit   = OrderType("STOP_LOSS_LIMIT")
	TypeTakeProfit      = OrderType("TAKE_PROFIT")
	TypeTakeProfitLimit = OrderType("TAKE_PROFIT_LIMIT")
	TypeLimitMaker      = OrderType("LIMIT_MAKER")

	SideBuy  = OrderSide("BUY")
	SideSell = OrderSide("SELL")
)
