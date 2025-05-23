package html_util

import (
	"context"
	"testing"
)

func TestParseHtml2Text(t *testing.T) {
	htmlStr := `<a name=\"ask\" onclick=\"return reask('160级武器种类');\" href=\"#\" class=\"sprite_color_link\">160级武器种类</a><br>　　<span class=\"sprite_color_text\">属性范围</span><span class=\"sprite_color_text\">：<br>　　</span><span class=\"sprite_color_text\">强化打造</span><span class=\"sprite_color_text\">：命中(571-777)；伤害(490-667)<br>　　</span><span class=\"sprite_color_text\">获得</span><span class=\"sprite_color_text\">：<br>　　1、用<a name=\"ask\" onclick=\"return reask('陨铁');\" href=\"#\" class=\"sprite_color_link\">陨铁</a></span>幻化150级的强化武器(有几率失败)，得到<a name=\"ask\" onclick=\"return reask('元身');\" href=\"#\" class=\"sprite_color_link\">元身</a><br>　　2、用元身与<a name=\"ask\" onclick=\"return reask('战魄');\" href=\"#\" class=\"sprite_color_link\">战魄</a>进行打造，触发任务，完成后找袁天罡(长安城357，245)领取未鉴定武器<br>　　<span class=\"sprite_color_text\">其他说明</span><span class=\"sprite_color_text\">：<br>　　1、幻化150级武器的属性不影响元身属性(建议用</span><span class=\"sprite_color_text\">150级白板国标</span><span class=\"sprite_color_text\">武器幻化)<br>　　2、失败3次的武器无法幻化<br>　　3、角色≥120级才能使用元身和战魄打造160装备</span><span class=\"sprite_color_text\"></span>`

	ctx := context.TODO()
	text, err := ParseHtml2Text(ctx, htmlStr)
	if err != nil {
		t.Error(err)
	}
	t.Log(text)
}
