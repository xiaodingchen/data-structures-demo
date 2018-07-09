<?php 


class UTF8Util
{
    public static function get_chars($str)
    {
        $s = $str;
        $len = strlen($s);
        if($len == 0) return [];
        // 将字符串转成utf8字符集
        $chars = [];
        for($i=0; $i<$len; $i++){
            $c = $s[$i];
            $n = ord($c);

            if(($n >> 7) == 0){       //0xxx xxxx, asci, single
                $chars[] = $c;
            }

            else if(($n >> 4) == 15){     //1111 xxxx, first in four char
                if($i < $len - 3){
                    $chars[] = $c.$s[$i + 1].$s[$i + 2].$s[$i + 3];
                    $i += 3;
                }
            }

            else if(($n >> 5) == 7){  //111x xxxx, first in three char
                if($i < $len - 2){
                    $chars[] = $c.$s[$i + 1].$s[$i + 2];
                    $i += 2;
                }
            }

            else if(($n >> 6) == 3){  //11xx xxxx, first in two char
                if($i < $len - 1){
                    $chars[] = $c.$s[$i + 1];
                    $i++;
                }
            }
        }

        return $chars;
    }
}


// 构造字典树(Tire tree)
class TrieTree
{
    private $tree = [];

    // 新增树节点
    public function insert($str)
    {
        $chars = UTF8Util::get_chars($str);
        $chars[] = null;
        $count = count($chars);
        $T = &$this->tree;
        for($i=0; $i<$count; $i++){
            $c = $chars[$i];
            if(!array_key_exists($c, $T)){
                $T[$c] = [];
            }

            $T = &$T[$c];
        }
    }

    // 删除一个子字典树

    public function remove($str)
    {
        $chars = UTF8Util::get_chars($str);
        $chars[] = null;
        // 先查找子树是否存在
        if(!$this->_find($chars)){
            return;
        }

        $count = count($chars);
        $T = &$this->tree;
        for($i=0; $i<$count; $i++){
            $c = $chars[$i];
            if(array_key_exists($c, $T)){
                if(count($T[$c]) == 1){
                    unset($T[$c]);
                    return;
                }
                $T = &$T[$c];
            }
        }

    }

    // 查找一个字符串
    public function find($str)
    {
        $chars = UTF8Util::get_chars($str);
        $chars[] = null;

        return $this->_find($chars);
    }

    /* 
     * 统计一个字符串中是否有字典树中的内容
     * @param string $str 待匹配的字符串
     * @param boolean $count 是否统计出现次数
     *
     * @return bool|int
     */
    public function contain($str, $do_count = false)
    {
        $chars = UTF8Util::get_chars($str);
        $chars[] = null;
        $len = count($chars);
        $count = 0;
        $Tree = &$this->tree;
        // 全文检索
        for($i=0; $i<$len; $i++){
            $c = $chars[$i];
            // 匹配起始字符(先确定起始字符)
            if(array_key_exists($c, $Tree)){
                $T = &$Tree[$c];
                for($j = $i+1; $j<$len; $j++){
                    $c = $chars[$j];
                    // 统计这个起始字符组成的字符串是否在字典树中
                    if(array_key_exists(null, $T)){
                        if($do_count){
                            $count++;
                        }
                        else{
                            return true;
                        }
                    }

                    if(!array_key_exists($c, $T)){
                        break;
                    }


                    $T = &$T[$c];
                }
            }
        }

        if($do_count){
            return $count;
        }

        return false;
    }

    public function contain_replace($str, $replace = '*')
    {
        $chars = UTF8Util::get_chars($str);
        $chars[] = null;
        $len = count($chars);
        $Tree = &$this->tree;
        // 全文检索
        for($i=0; $i<$len; $i++){
            $c = $chars[$i];
            // 匹配起始字符(先确定起始字符)
            if(array_key_exists($c, $Tree)){
                $T = &$Tree[$c];
                for($j = $i+1; $j<$len; $j++){
                    $c = $chars[$j];
                    // 统计这个起始字符组成的字符串是否在字典树中
                    if(array_key_exists(null, $T)){
                        $chars[$i] = $replace;
                        
                    }

                    if(!array_key_exists($c, $T)){
                        break;
                    }else{
                        $chars[$j] = $replace;
                    }

                    $T = &$T[$c];
                }
            }
        }

        return implode('', $chars);
    }

    private function _find($chars)
    {
        $count = count($chars);
        $T = &$this->tree;
        for($i=0; $i<$count; $i++){
            $c = $chars[$i];
            if(!array_key_exists($c, $T)){
                return false;
            }

            $T = &$T[$c];
        }

        return true;
    }

    public function getTree()
    {
        return $this->tree;
    }
}

// $str = '你好lll';
// $str2 = '你不231';
$tree = new TrieTree;


$minWord = [
    '国军',
    '江长生',
    '大新闻',
    '江长者',
    '大长明',
];

foreach ($minWord as $value) {
    $tree->insert($value);
}

$text = "经过大选，会议决定江长者当选为国军总司令";

print_r($tree->contain_replace($text));
print_r("\n");
print_r($tree->contain($text, true));

// var_dump(UTF8Util::get_chars($str));
