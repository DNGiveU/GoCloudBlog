package model

/**
 * ant design pro 键值数据结构
 */
type KLabel struct {
	Key string		`json:"key"`
	Label string	`json:"label"`
}

/**
 * 地理位置数据结构
 */
type Geographic struct {
	Province KLabel		`json:"province"`
	City KLabel			`json:"city"`
}