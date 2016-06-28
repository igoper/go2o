/**
 * Copyright 2015 @ z3q.net.
 * name : snapshot
 * author : jarryliu
 * date : 2016-06-28 21:41
 * description :
 * history :
 */
package goods

type (
	// 商品快照
	GoodsSnapshot struct {
		Id           int    `db:"id" auto:"yes" pk:"yes"`
		Key          string `db:"snapshot_key"`
		ItemId       int    `db:"item_id"`
		GoodsId      int    `db:"goods_id"`
		GoodsName    string `db:"goods_name"`
		GoodsNo      string `db:"goods_no"`
		SmallTitle   string `db:"small_title"`
		CategoryName string `db:"category_name"`
		Image        string `db:"img"`

		//成本价
		Cost float32 `db:"cost"`

		//定价
		Price float32 `db:"price"`

		//销售价
		SalePrice  float32 `db:"sale_price"`
		CreateTime int64   `db:"create_time"`
	}
)