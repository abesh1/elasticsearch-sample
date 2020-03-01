package mecab

// #include <mecab.h>
// #include <stdlib.h>
import "C"
import "strings"

type Tagger struct {
	tagger *C.mecab_t
}

func (tg *Tagger) Destroy() {
	C.mecab_destroy(tg.tagger)
}

func (tg *Tagger) Parse(lt *Lattice) string {
	C.mecab_parse_lattice(tg.tagger, lt.lattice)
	return C.GoString(C.mecab_lattice_tostr(lt.lattice))
}

func (tg *Tagger) NBestParse(lt *Lattice, n int) []string {
	C.mecab_lattice_set_request_type(lt.lattice, C.MECAB_NBEST)
	C.mecab_parse_lattice(tg.tagger, lt.lattice)

	_n := C.size_t(n)

	res := C.GoString(C.mecab_lattice_nbest_tostr(lt.lattice, _n))
	a := strings.Split(res, "\n")

	n2 := n
	if len(a) < n {
		n2 = len(a)
	}
	return a[:n2]
}

func (tg *Tagger) ParseToNode(lt *Lattice) *Node {
	C.mecab_parse_lattice(tg.tagger, lt.lattice)
	node := C.mecab_lattice_get_bos_node(lt.lattice)
	return &Node{node, node}
}
