package day21

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)\ntrh fvjkl sbzzf mxmxvkd (contains dairy)\nsqjhc fvjkl (contains soy)\nsqjhc mxmxvkd sbzzf (contains fish)",
			5},
		{
			"hhbf zkzls sbkb krd msmnzlm pbfnr szpbd dgbz mjrrfzz dzrjg mttl slgdh nmgqt hqkxv nchrg dvknd fpcfdg jtrgt nrrll cbmbhft rjnmg dzrzc vgfkcnr mtdsv tbthrm hdsm dml xchzh kfgd gzlqd khhf sjhds zk jxbqcp mzvkx nbgklf qlqrm rkjjsl fvszcb mzj kncpfb pdkcrm gvnm vvkrmf fttbhdr qxgtsm qjxxpr kmr xbfht kbrr bxvkx clvr rmblkb dgdq shqqg rldlxf qqfnqb gzbhn vzj gtqqlgl hnctrp (contains eggs, fish)\nxchzh vmn gcvqnz hkhd fttbhdr nbzzpx srk kmr gzbhn jtrgt zk rjnmg vlpzjs mznbn slgdh mmzg hdbjmn plqvl vqd nkhrnpf kbrr ltzrz fghm kfgd jslmtf qxgtsm lxdjq bthf rgtb pdkv nfnfk pjrlld vgsgnkq kjxlsmb jzqtjc pqxvmx ghrxm mttl jqtlkmz xfn zsbc jcfmh hfxxf fpcfdg kgdqhg lgjjh sbkb qlqrm qjxxpr bxvkx sjhds zlzkh sgpncx lrbbmj dgbz kncpfb jpljgr mjrrfzz zkzls nzkr nbgklf hqkxv gjssscmd jxbqcp hpvfng nchrg trq nlszr gpscr czpx dgdq qdlczn ltphvb vvkrmf hdsm xhpnq khtb (contains soy, dairy, wheat)\nxfn mznbn sgxzx mzvkx czpx grgj hdsm nbgklf hdjqhg lqvk lzjhp zvlpx vmn rldlxf nkhrnpf gxsfth ndpgv rfpnz jxbqcp xchzh nblsxmr pdkcrm kbrt jqtlkmz kzmbxzg tbthrm nvtmrs qxgtsm fttbhdr dtp fncgkl gjssscmd rmblkb sjkcx vhf rlrdk clvr rgffc kjxlsmb qqfnqb mxxmpfs mtdsv nfnfk qjxxpr gpscr fvszcb rrnt svjsqv bthf vhgt zsbc bnmttd lgjjh ltzrz ghrxm qgcxhbm kfgd mxzbbh kmdgt xcxvvx zlzkh jpljgr khhf (contains fish, nuts, sesame)\nqqfnqb dzrzc sjhds zdd hdsm fghm stdq hsxj qjxxpr dxhbtfmk dgpg jpxs nfnfk hnctrp plnf dtp kmr xcxvvx mlglz cszhgf nbgklf dgdq jpljgr gcvqnz gzlqd npmf tbthrm dntjn sgxzx kmdgt gpscr pptnq zk clvr vkvg mzvkx hfxxf lzjhp kncpfb qdgnkl mzj dvxq vkxsv bqhjlt htrzq stq xdnbq bxvkx svjsqv rfpnz zdmk pqxvmx hxcshv xchzh mntmxzf mttl ltphvb mmzg msmnzlm jqtlkmz kgdqhg nchrg lxdjq (contains soy)\nkkbsc kjxlsmb zhzspb kmr xcxvvx bsnnnhq nlszr qgcxhbm hdsm qjxxpr chqbxcd rldlxf nnpz vkvg pgxkln khhf nblsxmr gzlqd plnf jxbqcp dml vcgtt mttl nfnfk djlpcc fttbhdr gxsfth fvszcb mxzbbh zlzkh rgtb xhpnq tbthrm nmgqt dzrjg rdkx slgdh gpscr qdgnkl hnctrp pptnq bshd hxgl jqtlkmz dgdq nbgklf jpljgr rrnt lrbbmj jslmtf gtc njcqb gvnm sjhds xctm clvr rfpnz gcvqnz svjsqv (contains eggs)\ndtp jzqtjc gzbhn pdkv xbfht mttl rfpnz mznbn jqtlkmz gpscr lskhfb lxdjq dzrzc xchzh fflp kjxlsmb zhzspb zdd vmn njcqb vcgtt mlglz mxzbbh hsxj xktk fxzj dlpnp qjxxpr vvkrmf szpbd nblsxmr clvr zdmk gvnm sjhds chqbxcd hdsm zlzkh bqhjlt fpcfdg kkbsc mmzg pjrlld mxxmpfs nfnfk kgdqhg djlpcc bhkzzl hdjqhg nchrg dntjn nnpz qdgnkl ndpgv plqvl hfxxf zsbc vptd gjssscmd nbgklf htrzq hxcshv (contains peanuts, fish)\nnvtmrs plqvl jzqtjc sjhds jtrgt kmdgt qqfnqb stdq mlglz nblsxmr vkxsv fttbhdr lgjjh pdkv qlqrm rfpnz vlpzjs vgz mntmxzf kmhrsh vgfkcnr krd hfxxf bsnnnhq qjxxpr lzjhp zlzkh dgdq chqbxcd mmzg bhkzzl clvr hhbf dvxq dxhbtfmk qlkhg xdnbq vgsgnkq mzvkx nlszr vqd plnf dzrjg fpcfdg dntjn xchzh njcqb hkhd mzj kbrr fxzj nkgd rgffc hxgl kgdqhg rldlxf dgpg hdsm gcvqnz npmf vcgtt jcfmh ltphvb nmgqt vzj hpvfng kbrt pfcvnq sgpncx chxrpb vhf pgxkln gtqqlgl ljl vptd qdgnkl vmn nfnfk mtdsv msmnzlm fncgkl (contains wheat)\nfttbhdr vhgt nfnfk mjrrfzz pfcvnq bsnnnhq stdq njcqb rgffc nbgklf gjssscmd kbrt dgbz ljl bshd fncgkl nzkr krd fflp chqbxcd trq htrzq fxzj vzj vkxsv chxrpb hhbf dml qxgtsm gzlqd xcxvvx vlpzjs lzfb qdlczn kmr lqvk clvr zsbc hpvfng kmdgt rgtb kgdqhg gxsfth sjhds hdsm svjsqv xdnbq zhzspb pjrlld kzmbxzg rkjjsl jpxs xchzh xktk grgj dzrjg bthf qlkhg hqkxv gpscr rldlxf zk pgxkln mzj qgcxhbm (contains fish, nuts, sesame)\npqxvmx dzrzc khtb qjxxpr dzrjg xbfht srk qdgnkl njcqb hkhd gxsfth lskhfb mzj xchzh fncgkl mjrrfzz hxgl pdkv sjhds krd mntmxzf fttbhdr mxzbbh hdbjmn mmzg kkbsc mznbn bqhjlt nfnfk hsxj xlrtg hdsm djlpcc hqkxv nbgklf stdq dvknd gsmbm nbzzpx nrrll kmr vlpzjs sbkb jpxs czpx fvszcb bhkzzl kjxlsmb tvfshn nblsxmr lzfb zdmk mzvkx (contains peanuts)\nnvtmrs svjsqv zsbc bsnnnhq jzqtjc lzfb xchzh lqvk sgpncx vzj gxsfth jslmtf pbfnr xktk xbfht msmnzlm dlpnp rlrdk lxdjq kmhrsh sjhds pjrlld kmr bthf hxcshv zvlpx vcgtt gzlqd kjxlsmb kncpfb gtqqlgl gvnm qjxxpr mtdsv plqvl fcrmdrf hdsm rjnmg mzj vgsgnkq xhpnq fflp jpljgr bqzmb fxzj bhkzzl nblsxmr gtc sbkb vmn gjssscmd kfgd clvr fmj hdjqhg nnpz ljl krd dvknd ltzrz nfnfk khhf vhf xfn dtp bqhjlt vkxsv qlkhg dvxq kmdgt vgfkcnr nbgklf fvszcb hpvfng zdmk nkgd czpx lrbbmj dml dgpg pgxkln rldlxf htrzq (contains sesame)\nsrk nfnfk shqqg qgcxhbm gsmbm bqzmb qqfnqb nvtmrs sjhds zsbc fcrmdrf jqtlkmz xfn nchrg trq bhkzzl fttbhdr rmblkb rldlxf rrnt xcxvvx xktk vzj mtdsv xhpnq fvszcb zlzkh tbthrm zdmk fghm nbgklf mntmxzf xchzh fpcfdg kmdgt gvnm pjrlld grgj qjxxpr mxzbbh fxzj zk xbfht pptnq hdsm sgpncx (contains soy)\nbthf zjt mlglz kfgd fttbhdr xphq czpx dvxq pptnq plqvl fpcfdg nvtmrs svjsqv pdkv qqfnqb slgdh mzvkx nfnfk dgbz lgjjh zkzls hfxxf pjrlld stq tvfshn xlrtg xchzh jqtlkmz kzmbxzg dzrjg nbgklf kgdqhg nchrg zlzkh bxvkx rkjjsl rldlxf nzkr szpbd gtc cbmbhft mttl kjxlsmb qjxxpr xktk stdq pdkcrm fvszcb hdsm zsbc kmr bqzmb rfpnz krd lqvk pbfnr lskhfb xcxvvx gcvqnz sjhds fncgkl cfhnhp rdkx gzbhn ltzrz ghrxm vzj (contains eggs, peanuts)\nzlzkh vhf fpcfdg vmn dzrzc vkxsv clvr lzfb nvtmrs chqbxcd stq xlrtg mntmxzf grgj vgfkcnr kmdgt trq rmblkb pqxvmx kbrt lqvk gtqqlgl xktk jcfmh sjhds kkbsc nbgklf tbthrm cszhgf srk vzj pptnq npmf fttbhdr nlszr qlqrm qjxxpr svjsqv xfn hqkxv mmzg pfcvnq hdbjmn hdsm plqvl cfhnhp mxzbbh fflp dvxq zhzspb sjkcx bsnnnhq fvszcb kmhrsh rjnmg sgpncx vcgtt dgdq bxvkx gxsfth ltphvb nfnfk nmgqt (contains fish)\nkfgd njcqb bsnnnhq dxhbtfmk zvlpx xdnbq djlpcc fflp pdkcrm lzjhp pptnq npmf rldlxf vkxsv szpbd hpvfng zsbc khhf dvknd mxzbbh lxdjq nlszr plqvl rmblkb tbthrm qdlczn gcvqnz mtdsv hdbjmn msmnzlm lqvk fttbhdr kbrt hdsm xbfht rgffc bthf xphq hdjqhg jzqtjc shqqg clvr bshd nzkr rjnmg nblsxmr htrzq khtb nfnfk mzj pgxkln nbzzpx xchzh sjkcx dzrzc vgz hqkxv vgfkcnr pdkv pjrlld sjhds jxbqcp vgsgnkq kjxlsmb hkhd kmdgt kmr qjxxpr hxcshv (contains fish, dairy)\nkbrt vgz fxzj rlrdk ltzrz krd xfn fpcfdg htrzq xchzh vkxsv msmnzlm vhgt jpljgr dxhbtfmk sjhds mntmxzf pjrlld nzkr gvnm szpbd pdkcrm hdsm hdjqhg jpxs nnpz zlzkh hnctrp xktk zhzspb lrbbmj gtqqlgl svjsqv bthf rrnt nmgqt qdgnkl mmzg mttl zk pptnq lxdjq cbmbhft dntjn qjxxpr zsbc hpvfng clvr nkhrnpf vkvg fttbhdr rkjjsl pbfnr hdbjmn lgjjh chqbxcd ljl lskhfb vgfkcnr nbgklf kjxlsmb nbzzpx khtb tbthrm kfgd (contains wheat, peanuts, nuts)\nhfxxf jtrgt kkbsc sjhds gpscr hqkxv msmnzlm kmr qjxxpr clvr qxgtsm hdsm mtdsv lzjhp kzmbxzg mzvkx jpxs mxxmpfs vlpzjs npmf qdlczn slgdh chqbxcd dzrzc gcvqnz hkhd nchrg qlkhg xbfht xhpnq zvlpx zdd plqvl xchzh pdkcrm gxsfth chxrpb nkhrnpf gjssscmd nfnfk pjrlld vptd vzj vkxsv rmblkb vhf ltphvb ljl kgdqhg gzlqd lzfb gzbhn vgfkcnr fttbhdr ndpgv htrzq kmhrsh xfn bsnnnhq (contains nuts, dairy)\ngcvqnz pjrlld dtp vhf plnf sbkb khhf nmgqt bshd vgz khtb rgtb kncpfb clvr srk qlqrm nchrg dgbz rldlxf fncgkl shqqg fttbhdr trq jpljgr qjxxpr fcrmdrf vvkrmf hdsm dntjn vhgt nrrll slgdh qgcxhbm jpxs mxzbbh rlrdk hxcshv nkhrnpf nbgklf kmhrsh svjsqv szpbd zdd vkvg nfnfk njcqb vgsgnkq jcfmh ltphvb zhzspb mttl hqkxv mzj mmzg mtdsv gsmbm chxrpb dxhbtfmk lgjjh vptd fpcfdg nnpz ghrxm xphq fmj xchzh xhpnq vmn hsxj kbrr nvtmrs plqvl hhbf hpvfng kjxlsmb pgxkln lzfb jqtlkmz (contains wheat, fish, eggs)\ngzbhn bhkzzl qlkhg kmr hpvfng hdsm ltphvb gvnm mntmxzf plqvl trq pdkcrm xctm dgdq rrnt nbgklf qjxxpr vhf qgcxhbm lrbbmj zdd sjhds hxgl svjsqv xdnbq nfnfk bqhjlt dvxq jpxs gpscr hdbjmn stdq shqqg clvr szpbd mzj jslmtf fttbhdr (contains fish)\nxktk bnmttd vkvg kncpfb gcvqnz dml dgdq hsxj sjhds nlszr pdkv gzbhn sbkb zvlpx hdsm nzkr vvkrmf mlglz chxrpb kbrr qjxxpr bxvkx zjt plnf qlqrm bqhjlt hnctrp khhf vptd dvxq kmr hdbjmn stdq fttbhdr vkxsv fvszcb ltphvb jpxs zkzls plqvl nvtmrs nfnfk nbzzpx pdkcrm cfhnhp nbgklf mzvkx vqd zk clvr kzmbxzg tvfshn stq jxbqcp trq bthf lrbbmj (contains soy)\nnblsxmr clvr sjhds vvkrmf qjxxpr nfnfk zkzls qgcxhbm rldlxf jpljgr qlkhg nbgklf rfpnz kjxlsmb qxgtsm ltzrz fttbhdr mtdsv pptnq bthf xlrtg nnpz hqkxv fncgkl hsxj chxrpb cfhnhp kncpfb fflp hxcshv mxxmpfs nzkr ltphvb xcxvvx qlqrm bvxkvt lqvk pdkcrm nkgd vgz mzvkx plqvl zsbc pbfnr bhkzzl jxbqcp mjrrfzz gcvqnz hfxxf gpscr gzlqd kfgd pfcvnq vlpzjs hdsm ljl dzrzc hhbf dgbz dvknd nkhrnpf mznbn khhf xhpnq lgjjh ndpgv jcfmh fvszcb zdmk mmzg (contains soy)\nszpbd gjssscmd qxgtsm sjhds chxrpb nmgqt xphq qdlczn jtrgt mttl sgxzx rgffc rrnt bhkzzl hqkxv zsbc nbgklf xctm nlszr vvkrmf fttbhdr sgpncx bvxkvt jcfmh htrzq sbkb dlpnp mntmxzf hdsm vzj mxxmpfs nzkr rfpnz fflp dntjn rgtb xbfht kzmbxzg kjxlsmb pptnq rlrdk tvfshn sjkcx clvr kmr npmf jpxs vmn vlpzjs cfhnhp pqxvmx pgxkln dxhbtfmk rkjjsl nfnfk nkgd rjnmg xchzh bshd rldlxf (contains sesame)\nzdmk qgcxhbm fncgkl jzqtjc trq rgtb gtqqlgl vmn clvr dzrzc xbfht hnctrp zkzls nbgklf zk qqfnqb lrbbmj srk kkbsc zhzspb rjnmg fxzj rdkx vptd xcxvvx stq vkvg dzrjg pgxkln rmblkb rfpnz gtc fcrmdrf xchzh kmr sjhds kjxlsmb xphq fflp nkhrnpf fttbhdr bqhjlt kzmbxzg lzjhp bhkzzl lxdjq xfn mtdsv hdsm xlrtg vgz hsxj plqvl fmj nbzzpx vcgtt gvnm bshd jxbqcp dtp fghm kmhrsh mxxmpfs ljl zdd chxrpb nfnfk gpscr hfxxf dgbz mmzg hxgl mxzbbh zjt (contains nuts, dairy, sesame)\nbshd xktk xphq rfpnz shqqg vvkrmf tbthrm srk dgbz kgdqhg mzj vcgtt sjkcx hxgl dgdq jzqtjc pdkv fxzj xhpnq gsmbm vlpzjs hsxj qgcxhbm nbzzpx tvfshn pgxkln jslmtf vgz kjxlsmb nfnfk sgpncx clvr rjnmg qjxxpr jqtlkmz vkvg dxhbtfmk fttbhdr xfn xchzh nbgklf krd rrnt jpxs mntmxzf dvxq kbrt kmr ltzrz mlglz fmj svjsqv slgdh nlszr dzrzc zvlpx qdlczn ndpgv fcrmdrf lzjhp dcqp sjhds kkbsc (contains dairy)\nsrk kgdqhg hqkxv fvszcb kkbsc pjrlld nrrll njcqb clvr nchrg jxbqcp kncpfb bnmttd fttbhdr xfn vqd rfpnz kmr xchzh vptd gzbhn svjsqv dntjn dxhbtfmk mxxmpfs rdkx hdjqhg mzvkx hdsm zsbc lzjhp hfxxf dgbz pfcvnq zjt zdmk fcrmdrf dzrjg rmblkb dlpnp hsxj qjxxpr ltzrz qgcxhbm vlpzjs rrnt bqzmb jzqtjc sbkb mlglz xhpnq xdnbq kfgd vgfkcnr vzj vkvg zdd htrzq tvfshn pdkv grgj jpxs slgdh mntmxzf xktk sjkcx bqhjlt sjhds nbgklf bxvkx jtrgt (contains wheat, peanuts, soy)\nhdsm nfnfk gpscr ltzrz rgtb pdkcrm chxrpb vkvg mttl jxbqcp vqd nlszr xbfht vhgt fxzj npmf rmblkb kkbsc pfcvnq xcxvvx pbfnr gxsfth xchzh qxgtsm slgdh pqxvmx zdmk vkxsv khhf hdjqhg rrnt nrrll nbgklf rlrdk zjt rfpnz bqzmb gvnm mxxmpfs czpx ljl vcgtt vhf rdkx dvxq xhpnq hsxj chqbxcd njcqb fpcfdg shqqg mxzbbh vgz stdq htrzq dxhbtfmk fcrmdrf hhbf jslmtf dntjn pptnq tvfshn zkzls sjhds vgsgnkq fmj lxdjq pgxkln jtrgt mtdsv gsmbm bhkzzl tbthrm qdgnkl rkjjsl xktk lskhfb clvr jpljgr dgbz jpxs bsnnnhq lrbbmj nblsxmr fttbhdr sjkcx (contains eggs)\nltphvb vptd nnpz rdkx dvxq srk tvfshn xctm xchzh clvr vgfkcnr hdsm lxdjq gsmbm szpbd hnctrp bxvkx msmnzlm pptnq zvlpx hfxxf qlkhg svjsqv fttbhdr rldlxf jzqtjc lgjjh nchrg vlpzjs kmdgt kmhrsh dgbz fpcfdg nbzzpx vhgt nkgd kkbsc bvxkvt vvkrmf nrrll pfcvnq qlqrm fncgkl rrnt nfnfk vhf hqkxv kjxlsmb sjhds qjxxpr hxcshv gzlqd pqxvmx bthf qdgnkl htrzq sgxzx nlszr lskhfb (contains soy, peanuts, sesame)\nhdjqhg njcqb dgbz ltphvb nblsxmr kfgd nbzzpx sbkb nmgqt jqtlkmz pdkcrm rkjjsl jpljgr sjhds fttbhdr vlpzjs hfxxf jcfmh chxrpb sjkcx cbmbhft dgdq gzbhn kmdgt trq xfn zlzkh nzkr pjrlld nfnfk dntjn jslmtf kjxlsmb zkzls sgpncx bxvkx srk nlszr xchzh npmf mzvkx gsmbm lskhfb clvr hdsm vhgt ljl stq hhbf dlpnp qjxxpr mznbn bthf fpcfdg dvknd plqvl fncgkl lqvk mzj qdgnkl gvnm kgdqhg hnctrp tbthrm (contains dairy, fish, peanuts)\nzkzls nblsxmr bhkzzl ltzrz hqkxv pgxkln hxgl dxhbtfmk qjxxpr gcvqnz khtb gpscr dgpg nvtmrs stdq mlglz srk xphq xhpnq kkbsc zk bthf vlpzjs chqbxcd lzjhp vptd vzj nfnfk nbgklf xlrtg fttbhdr shqqg vhgt rldlxf mxzbbh bxvkx xdnbq xchzh hnctrp vgsgnkq rgtb zhzspb sjhds xbfht pdkcrm gjssscmd dzrzc mzj zjt hdsm tbthrm pqxvmx jtrgt jpljgr czpx gzbhn dvxq lqvk vmn lgjjh qqfnqb hdbjmn qdgnkl fflp jcfmh hdjqhg vvkrmf (contains dairy)\ngrgj zkzls xbfht dtp rjnmg zsbc nfnfk sgxzx jtrgt hdsm hhbf vmn kfgd tbthrm slgdh gsmbm kjxlsmb nkhrnpf xdnbq gzlqd qjxxpr vkxsv vzj mjrrfzz hfxxf ndpgv pfcvnq krd cbmbhft sjhds djlpcc clvr lrbbmj pptnq bqhjlt rgtb zhzspb kmhrsh vgz nzkr xchzh stdq sgpncx vlpzjs bthf jslmtf gcvqnz pgxkln hkhd mzvkx fvszcb gjssscmd lqvk shqqg dgdq qlqrm njcqb gtqqlgl qlkhg qxgtsm plqvl zk vhf fttbhdr khtb zlzkh rmblkb (contains peanuts, fish)\nclvr vcgtt lzfb kkbsc nfnfk lzjhp lgjjh trq bqzmb stdq xlrtg fxzj qqfnqb gvnm vkvg hxcshv kmr jzqtjc gtc ghrxm mmzg zdmk fflp krd vkxsv fncgkl jpxs dgbz mntmxzf fttbhdr zk lskhfb jcfmh dntjn hpvfng jqtlkmz qjxxpr zdd pqxvmx kgdqhg khhf kbrr mznbn jxbqcp zjt nnpz rgffc mxzbbh mttl qgcxhbm nbgklf sgpncx dvxq pbfnr xhpnq fcrmdrf szpbd dlpnp rldlxf nrrll shqqg pdkv sjhds khtb rfpnz xchzh nkgd xctm rdkx (contains dairy, nuts, eggs)\nzsbc nfnfk htrzq jpxs qxgtsm dgbz khhf hdjqhg fmj nlszr tbthrm qjxxpr dxhbtfmk rrnt gpscr mznbn nbgklf kncpfb dgdq fcrmdrf qlqrm jtrgt gtqqlgl clvr nchrg rgtb hdsm xchzh sgxzx fttbhdr zlzkh cszhgf xphq ltphvb pdkv cbmbhft djlpcc vgsgnkq xfn rldlxf msmnzlm vptd vhgt (contains nuts, wheat, sesame)\nnjcqb clvr pptnq vcgtt kmr xfn vlpzjs mznbn vhgt jxbqcp kfgd dvknd qlkhg nrrll fttbhdr hdsm bhkzzl svjsqv fcrmdrf xlrtg nfnfk hdjqhg stq nzkr msmnzlm rmblkb grgj mxxmpfs rlrdk khtb rfpnz dml fflp kbrr sjhds qjxxpr hfxxf jzqtjc gjssscmd mzj cszhgf szpbd dntjn nbgklf jslmtf qlqrm (contains soy, sesame)\nvkxsv zsbc mxzbbh nvtmrs plqvl pdkv mtdsv bsnnnhq hdsm pjrlld hxgl zdd kmdgt xdnbq kkbsc qxgtsm qdlczn qqfnqb vzj dntjn pdkcrm nfnfk fcrmdrf djlpcc hxcshv njcqb jslmtf lxdjq srk clvr cszhgf nblsxmr dgbz sbkb rldlxf lrbbmj lqvk ghrxm kbrr rgffc gzlqd tbthrm hsxj stdq sjhds nlszr nmgqt vmn vgfkcnr nzkr vkvg fttbhdr rrnt lzfb zvlpx xchzh fvszcb chxrpb fflp zk dzrjg gcvqnz jzqtjc hpvfng lskhfb rmblkb tvfshn pqxvmx vvkrmf nkhrnpf zlzkh svjsqv nbgklf fghm htrzq sgxzx vcgtt mmzg pptnq vlpzjs hfxxf dcqp (contains soy, nuts, sesame)\nfxzj stdq nmgqt hkhd sgxzx tvfshn clvr jzqtjc gcvqnz nvtmrs rmblkb kkbsc bnmttd nrrll slgdh cszhgf nkgd chqbxcd hpvfng pdkv lskhfb gtqqlgl hxgl xphq bqzmb bsnnnhq qjxxpr mxxmpfs zjt nblsxmr lzfb zdd jtrgt xcxvvx nbgklf mzvkx dtp nfnfk mxzbbh mttl hhbf gzbhn rdkx bhkzzl vgsgnkq vgfkcnr hdbjmn jslmtf krd sjhds vhgt qlqrm zdmk srk nnpz hqkxv zhzspb vkvg bvxkvt xchzh hsxj fttbhdr czpx ljl vhf rfpnz bqhjlt gvnm dcqp hnctrp mtdsv mjrrfzz qgcxhbm mznbn (contains fish)\nkbrt kncpfb vhf npmf kmr kbrr htrzq rldlxf xhpnq plqvl qgcxhbm mjrrfzz vvkrmf zdmk xctm gtc jxbqcp xktk hqkxv mlglz fpcfdg hkhd kgdqhg jzqtjc rfpnz shqqg grgj zk sgpncx gjssscmd qxgtsm czpx vhgt zkzls sgxzx xchzh lgjjh pjrlld djlpcc qdgnkl pgxkln fcrmdrf nlszr ltphvb ljl qlqrm xbfht hdsm bvxkvt clvr pqxvmx qjxxpr nkhrnpf fghm dlpnp vptd srk gsmbm zsbc hfxxf fttbhdr cbmbhft xfn xdnbq rdkx sjhds stq dgpg nbgklf (contains eggs, fish, sesame)\nxphq vgsgnkq kkbsc kjxlsmb stq qlqrm jpxs mmzg lxdjq xlrtg nchrg dxhbtfmk ljl ndpgv khhf rdkx qlkhg vmn lrbbmj gtc krd pfcvnq rgffc dntjn jslmtf dtp mttl lskhfb rfpnz jpljgr hnctrp fxzj rldlxf kfgd mznbn bshd hkhd gxsfth kbrt fflp lzjhp fttbhdr msmnzlm xchzh kncpfb sjhds vzj bqhjlt mxxmpfs svjsqv qjxxpr nbzzpx jxbqcp rrnt hdjqhg cbmbhft gvnm sbkb bhkzzl rmblkb qxgtsm stdq clvr gcvqnz hhbf hdsm nfnfk vgz bvxkvt rlrdk jzqtjc hpvfng nrrll zsbc mjrrfzz pdkcrm (contains nuts)\nvkxsv gzlqd zsbc htrzq qlqrm dlpnp jzqtjc hdsm ghrxm dvknd hfxxf vcgtt hkhd dcqp vgfkcnr zk rrnt svjsqv nblsxmr qjxxpr xcxvvx qdgnkl ltzrz lzjhp hxgl kkbsc mttl nfnfk ljl rgffc stq shqqg jtrgt nkhrnpf msmnzlm dtp mmzg hpvfng fmj nnpz vlpzjs vqd rlrdk mxzbbh mntmxzf rkjjsl cfhnhp zkzls sgxzx sjhds kbrt hqkxv jcfmh nmgqt fttbhdr nzkr qlkhg xctm szpbd stdq lqvk nlszr krd mjrrfzz clvr trq grgj hhbf vmn pptnq qxgtsm nbgklf dntjn nrrll lgjjh gtqqlgl dgpg xdnbq fcrmdrf djlpcc (contains nuts, eggs, peanuts)\n",
			2170},
	}

	for i, test := range tests {
		actual := part1(test.in)
		if actual != test.want {
			t.Errorf("#%d (part1): actual = %d, but want %d", i, actual, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)\ntrh fvjkl sbzzf mxmxvkd (contains dairy)\nsqjhc fvjkl (contains soy)\nsqjhc mxmxvkd sbzzf (contains fish)",
			"mxmxvkd,sqjhc,fvjkl"},
		{
			"hhbf zkzls sbkb krd msmnzlm pbfnr szpbd dgbz mjrrfzz dzrjg mttl slgdh nmgqt hqkxv nchrg dvknd fpcfdg jtrgt nrrll cbmbhft rjnmg dzrzc vgfkcnr mtdsv tbthrm hdsm dml xchzh kfgd gzlqd khhf sjhds zk jxbqcp mzvkx nbgklf qlqrm rkjjsl fvszcb mzj kncpfb pdkcrm gvnm vvkrmf fttbhdr qxgtsm qjxxpr kmr xbfht kbrr bxvkx clvr rmblkb dgdq shqqg rldlxf qqfnqb gzbhn vzj gtqqlgl hnctrp (contains eggs, fish)\nxchzh vmn gcvqnz hkhd fttbhdr nbzzpx srk kmr gzbhn jtrgt zk rjnmg vlpzjs mznbn slgdh mmzg hdbjmn plqvl vqd nkhrnpf kbrr ltzrz fghm kfgd jslmtf qxgtsm lxdjq bthf rgtb pdkv nfnfk pjrlld vgsgnkq kjxlsmb jzqtjc pqxvmx ghrxm mttl jqtlkmz xfn zsbc jcfmh hfxxf fpcfdg kgdqhg lgjjh sbkb qlqrm qjxxpr bxvkx sjhds zlzkh sgpncx lrbbmj dgbz kncpfb jpljgr mjrrfzz zkzls nzkr nbgklf hqkxv gjssscmd jxbqcp hpvfng nchrg trq nlszr gpscr czpx dgdq qdlczn ltphvb vvkrmf hdsm xhpnq khtb (contains soy, dairy, wheat)\nxfn mznbn sgxzx mzvkx czpx grgj hdsm nbgklf hdjqhg lqvk lzjhp zvlpx vmn rldlxf nkhrnpf gxsfth ndpgv rfpnz jxbqcp xchzh nblsxmr pdkcrm kbrt jqtlkmz kzmbxzg tbthrm nvtmrs qxgtsm fttbhdr dtp fncgkl gjssscmd rmblkb sjkcx vhf rlrdk clvr rgffc kjxlsmb qqfnqb mxxmpfs mtdsv nfnfk qjxxpr gpscr fvszcb rrnt svjsqv bthf vhgt zsbc bnmttd lgjjh ltzrz ghrxm qgcxhbm kfgd mxzbbh kmdgt xcxvvx zlzkh jpljgr khhf (contains fish, nuts, sesame)\nqqfnqb dzrzc sjhds zdd hdsm fghm stdq hsxj qjxxpr dxhbtfmk dgpg jpxs nfnfk hnctrp plnf dtp kmr xcxvvx mlglz cszhgf nbgklf dgdq jpljgr gcvqnz gzlqd npmf tbthrm dntjn sgxzx kmdgt gpscr pptnq zk clvr vkvg mzvkx hfxxf lzjhp kncpfb qdgnkl mzj dvxq vkxsv bqhjlt htrzq stq xdnbq bxvkx svjsqv rfpnz zdmk pqxvmx hxcshv xchzh mntmxzf mttl ltphvb mmzg msmnzlm jqtlkmz kgdqhg nchrg lxdjq (contains soy)\nkkbsc kjxlsmb zhzspb kmr xcxvvx bsnnnhq nlszr qgcxhbm hdsm qjxxpr chqbxcd rldlxf nnpz vkvg pgxkln khhf nblsxmr gzlqd plnf jxbqcp dml vcgtt mttl nfnfk djlpcc fttbhdr gxsfth fvszcb mxzbbh zlzkh rgtb xhpnq tbthrm nmgqt dzrjg rdkx slgdh gpscr qdgnkl hnctrp pptnq bshd hxgl jqtlkmz dgdq nbgklf jpljgr rrnt lrbbmj jslmtf gtc njcqb gvnm sjhds xctm clvr rfpnz gcvqnz svjsqv (contains eggs)\ndtp jzqtjc gzbhn pdkv xbfht mttl rfpnz mznbn jqtlkmz gpscr lskhfb lxdjq dzrzc xchzh fflp kjxlsmb zhzspb zdd vmn njcqb vcgtt mlglz mxzbbh hsxj xktk fxzj dlpnp qjxxpr vvkrmf szpbd nblsxmr clvr zdmk gvnm sjhds chqbxcd hdsm zlzkh bqhjlt fpcfdg kkbsc mmzg pjrlld mxxmpfs nfnfk kgdqhg djlpcc bhkzzl hdjqhg nchrg dntjn nnpz qdgnkl ndpgv plqvl hfxxf zsbc vptd gjssscmd nbgklf htrzq hxcshv (contains peanuts, fish)\nnvtmrs plqvl jzqtjc sjhds jtrgt kmdgt qqfnqb stdq mlglz nblsxmr vkxsv fttbhdr lgjjh pdkv qlqrm rfpnz vlpzjs vgz mntmxzf kmhrsh vgfkcnr krd hfxxf bsnnnhq qjxxpr lzjhp zlzkh dgdq chqbxcd mmzg bhkzzl clvr hhbf dvxq dxhbtfmk qlkhg xdnbq vgsgnkq mzvkx nlszr vqd plnf dzrjg fpcfdg dntjn xchzh njcqb hkhd mzj kbrr fxzj nkgd rgffc hxgl kgdqhg rldlxf dgpg hdsm gcvqnz npmf vcgtt jcfmh ltphvb nmgqt vzj hpvfng kbrt pfcvnq sgpncx chxrpb vhf pgxkln gtqqlgl ljl vptd qdgnkl vmn nfnfk mtdsv msmnzlm fncgkl (contains wheat)\nfttbhdr vhgt nfnfk mjrrfzz pfcvnq bsnnnhq stdq njcqb rgffc nbgklf gjssscmd kbrt dgbz ljl bshd fncgkl nzkr krd fflp chqbxcd trq htrzq fxzj vzj vkxsv chxrpb hhbf dml qxgtsm gzlqd xcxvvx vlpzjs lzfb qdlczn kmr lqvk clvr zsbc hpvfng kmdgt rgtb kgdqhg gxsfth sjhds hdsm svjsqv xdnbq zhzspb pjrlld kzmbxzg rkjjsl jpxs xchzh xktk grgj dzrjg bthf qlkhg hqkxv gpscr rldlxf zk pgxkln mzj qgcxhbm (contains fish, nuts, sesame)\npqxvmx dzrzc khtb qjxxpr dzrjg xbfht srk qdgnkl njcqb hkhd gxsfth lskhfb mzj xchzh fncgkl mjrrfzz hxgl pdkv sjhds krd mntmxzf fttbhdr mxzbbh hdbjmn mmzg kkbsc mznbn bqhjlt nfnfk hsxj xlrtg hdsm djlpcc hqkxv nbgklf stdq dvknd gsmbm nbzzpx nrrll kmr vlpzjs sbkb jpxs czpx fvszcb bhkzzl kjxlsmb tvfshn nblsxmr lzfb zdmk mzvkx (contains peanuts)\nnvtmrs svjsqv zsbc bsnnnhq jzqtjc lzfb xchzh lqvk sgpncx vzj gxsfth jslmtf pbfnr xktk xbfht msmnzlm dlpnp rlrdk lxdjq kmhrsh sjhds pjrlld kmr bthf hxcshv zvlpx vcgtt gzlqd kjxlsmb kncpfb gtqqlgl gvnm qjxxpr mtdsv plqvl fcrmdrf hdsm rjnmg mzj vgsgnkq xhpnq fflp jpljgr bqzmb fxzj bhkzzl nblsxmr gtc sbkb vmn gjssscmd kfgd clvr fmj hdjqhg nnpz ljl krd dvknd ltzrz nfnfk khhf vhf xfn dtp bqhjlt vkxsv qlkhg dvxq kmdgt vgfkcnr nbgklf fvszcb hpvfng zdmk nkgd czpx lrbbmj dml dgpg pgxkln rldlxf htrzq (contains sesame)\nsrk nfnfk shqqg qgcxhbm gsmbm bqzmb qqfnqb nvtmrs sjhds zsbc fcrmdrf jqtlkmz xfn nchrg trq bhkzzl fttbhdr rmblkb rldlxf rrnt xcxvvx xktk vzj mtdsv xhpnq fvszcb zlzkh tbthrm zdmk fghm nbgklf mntmxzf xchzh fpcfdg kmdgt gvnm pjrlld grgj qjxxpr mxzbbh fxzj zk xbfht pptnq hdsm sgpncx (contains soy)\nbthf zjt mlglz kfgd fttbhdr xphq czpx dvxq pptnq plqvl fpcfdg nvtmrs svjsqv pdkv qqfnqb slgdh mzvkx nfnfk dgbz lgjjh zkzls hfxxf pjrlld stq tvfshn xlrtg xchzh jqtlkmz kzmbxzg dzrjg nbgklf kgdqhg nchrg zlzkh bxvkx rkjjsl rldlxf nzkr szpbd gtc cbmbhft mttl kjxlsmb qjxxpr xktk stdq pdkcrm fvszcb hdsm zsbc kmr bqzmb rfpnz krd lqvk pbfnr lskhfb xcxvvx gcvqnz sjhds fncgkl cfhnhp rdkx gzbhn ltzrz ghrxm vzj (contains eggs, peanuts)\nzlzkh vhf fpcfdg vmn dzrzc vkxsv clvr lzfb nvtmrs chqbxcd stq xlrtg mntmxzf grgj vgfkcnr kmdgt trq rmblkb pqxvmx kbrt lqvk gtqqlgl xktk jcfmh sjhds kkbsc nbgklf tbthrm cszhgf srk vzj pptnq npmf fttbhdr nlszr qlqrm qjxxpr svjsqv xfn hqkxv mmzg pfcvnq hdbjmn hdsm plqvl cfhnhp mxzbbh fflp dvxq zhzspb sjkcx bsnnnhq fvszcb kmhrsh rjnmg sgpncx vcgtt dgdq bxvkx gxsfth ltphvb nfnfk nmgqt (contains fish)\nkfgd njcqb bsnnnhq dxhbtfmk zvlpx xdnbq djlpcc fflp pdkcrm lzjhp pptnq npmf rldlxf vkxsv szpbd hpvfng zsbc khhf dvknd mxzbbh lxdjq nlszr plqvl rmblkb tbthrm qdlczn gcvqnz mtdsv hdbjmn msmnzlm lqvk fttbhdr kbrt hdsm xbfht rgffc bthf xphq hdjqhg jzqtjc shqqg clvr bshd nzkr rjnmg nblsxmr htrzq khtb nfnfk mzj pgxkln nbzzpx xchzh sjkcx dzrzc vgz hqkxv vgfkcnr pdkv pjrlld sjhds jxbqcp vgsgnkq kjxlsmb hkhd kmdgt kmr qjxxpr hxcshv (contains fish, dairy)\nkbrt vgz fxzj rlrdk ltzrz krd xfn fpcfdg htrzq xchzh vkxsv msmnzlm vhgt jpljgr dxhbtfmk sjhds mntmxzf pjrlld nzkr gvnm szpbd pdkcrm hdsm hdjqhg jpxs nnpz zlzkh hnctrp xktk zhzspb lrbbmj gtqqlgl svjsqv bthf rrnt nmgqt qdgnkl mmzg mttl zk pptnq lxdjq cbmbhft dntjn qjxxpr zsbc hpvfng clvr nkhrnpf vkvg fttbhdr rkjjsl pbfnr hdbjmn lgjjh chqbxcd ljl lskhfb vgfkcnr nbgklf kjxlsmb nbzzpx khtb tbthrm kfgd (contains wheat, peanuts, nuts)\nhfxxf jtrgt kkbsc sjhds gpscr hqkxv msmnzlm kmr qjxxpr clvr qxgtsm hdsm mtdsv lzjhp kzmbxzg mzvkx jpxs mxxmpfs vlpzjs npmf qdlczn slgdh chqbxcd dzrzc gcvqnz hkhd nchrg qlkhg xbfht xhpnq zvlpx zdd plqvl xchzh pdkcrm gxsfth chxrpb nkhrnpf gjssscmd nfnfk pjrlld vptd vzj vkxsv rmblkb vhf ltphvb ljl kgdqhg gzlqd lzfb gzbhn vgfkcnr fttbhdr ndpgv htrzq kmhrsh xfn bsnnnhq (contains nuts, dairy)\ngcvqnz pjrlld dtp vhf plnf sbkb khhf nmgqt bshd vgz khtb rgtb kncpfb clvr srk qlqrm nchrg dgbz rldlxf fncgkl shqqg fttbhdr trq jpljgr qjxxpr fcrmdrf vvkrmf hdsm dntjn vhgt nrrll slgdh qgcxhbm jpxs mxzbbh rlrdk hxcshv nkhrnpf nbgklf kmhrsh svjsqv szpbd zdd vkvg nfnfk njcqb vgsgnkq jcfmh ltphvb zhzspb mttl hqkxv mzj mmzg mtdsv gsmbm chxrpb dxhbtfmk lgjjh vptd fpcfdg nnpz ghrxm xphq fmj xchzh xhpnq vmn hsxj kbrr nvtmrs plqvl hhbf hpvfng kjxlsmb pgxkln lzfb jqtlkmz (contains wheat, fish, eggs)\ngzbhn bhkzzl qlkhg kmr hpvfng hdsm ltphvb gvnm mntmxzf plqvl trq pdkcrm xctm dgdq rrnt nbgklf qjxxpr vhf qgcxhbm lrbbmj zdd sjhds hxgl svjsqv xdnbq nfnfk bqhjlt dvxq jpxs gpscr hdbjmn stdq shqqg clvr szpbd mzj jslmtf fttbhdr (contains fish)\nxktk bnmttd vkvg kncpfb gcvqnz dml dgdq hsxj sjhds nlszr pdkv gzbhn sbkb zvlpx hdsm nzkr vvkrmf mlglz chxrpb kbrr qjxxpr bxvkx zjt plnf qlqrm bqhjlt hnctrp khhf vptd dvxq kmr hdbjmn stdq fttbhdr vkxsv fvszcb ltphvb jpxs zkzls plqvl nvtmrs nfnfk nbzzpx pdkcrm cfhnhp nbgklf mzvkx vqd zk clvr kzmbxzg tvfshn stq jxbqcp trq bthf lrbbmj (contains soy)\nnblsxmr clvr sjhds vvkrmf qjxxpr nfnfk zkzls qgcxhbm rldlxf jpljgr qlkhg nbgklf rfpnz kjxlsmb qxgtsm ltzrz fttbhdr mtdsv pptnq bthf xlrtg nnpz hqkxv fncgkl hsxj chxrpb cfhnhp kncpfb fflp hxcshv mxxmpfs nzkr ltphvb xcxvvx qlqrm bvxkvt lqvk pdkcrm nkgd vgz mzvkx plqvl zsbc pbfnr bhkzzl jxbqcp mjrrfzz gcvqnz hfxxf gpscr gzlqd kfgd pfcvnq vlpzjs hdsm ljl dzrzc hhbf dgbz dvknd nkhrnpf mznbn khhf xhpnq lgjjh ndpgv jcfmh fvszcb zdmk mmzg (contains soy)\nszpbd gjssscmd qxgtsm sjhds chxrpb nmgqt xphq qdlczn jtrgt mttl sgxzx rgffc rrnt bhkzzl hqkxv zsbc nbgklf xctm nlszr vvkrmf fttbhdr sgpncx bvxkvt jcfmh htrzq sbkb dlpnp mntmxzf hdsm vzj mxxmpfs nzkr rfpnz fflp dntjn rgtb xbfht kzmbxzg kjxlsmb pptnq rlrdk tvfshn sjkcx clvr kmr npmf jpxs vmn vlpzjs cfhnhp pqxvmx pgxkln dxhbtfmk rkjjsl nfnfk nkgd rjnmg xchzh bshd rldlxf (contains sesame)\nzdmk qgcxhbm fncgkl jzqtjc trq rgtb gtqqlgl vmn clvr dzrzc xbfht hnctrp zkzls nbgklf zk qqfnqb lrbbmj srk kkbsc zhzspb rjnmg fxzj rdkx vptd xcxvvx stq vkvg dzrjg pgxkln rmblkb rfpnz gtc fcrmdrf xchzh kmr sjhds kjxlsmb xphq fflp nkhrnpf fttbhdr bqhjlt kzmbxzg lzjhp bhkzzl lxdjq xfn mtdsv hdsm xlrtg vgz hsxj plqvl fmj nbzzpx vcgtt gvnm bshd jxbqcp dtp fghm kmhrsh mxxmpfs ljl zdd chxrpb nfnfk gpscr hfxxf dgbz mmzg hxgl mxzbbh zjt (contains nuts, dairy, sesame)\nbshd xktk xphq rfpnz shqqg vvkrmf tbthrm srk dgbz kgdqhg mzj vcgtt sjkcx hxgl dgdq jzqtjc pdkv fxzj xhpnq gsmbm vlpzjs hsxj qgcxhbm nbzzpx tvfshn pgxkln jslmtf vgz kjxlsmb nfnfk sgpncx clvr rjnmg qjxxpr jqtlkmz vkvg dxhbtfmk fttbhdr xfn xchzh nbgklf krd rrnt jpxs mntmxzf dvxq kbrt kmr ltzrz mlglz fmj svjsqv slgdh nlszr dzrzc zvlpx qdlczn ndpgv fcrmdrf lzjhp dcqp sjhds kkbsc (contains dairy)\nsrk kgdqhg hqkxv fvszcb kkbsc pjrlld nrrll njcqb clvr nchrg jxbqcp kncpfb bnmttd fttbhdr xfn vqd rfpnz kmr xchzh vptd gzbhn svjsqv dntjn dxhbtfmk mxxmpfs rdkx hdjqhg mzvkx hdsm zsbc lzjhp hfxxf dgbz pfcvnq zjt zdmk fcrmdrf dzrjg rmblkb dlpnp hsxj qjxxpr ltzrz qgcxhbm vlpzjs rrnt bqzmb jzqtjc sbkb mlglz xhpnq xdnbq kfgd vgfkcnr vzj vkvg zdd htrzq tvfshn pdkv grgj jpxs slgdh mntmxzf xktk sjkcx bqhjlt sjhds nbgklf bxvkx jtrgt (contains wheat, peanuts, soy)\nhdsm nfnfk gpscr ltzrz rgtb pdkcrm chxrpb vkvg mttl jxbqcp vqd nlszr xbfht vhgt fxzj npmf rmblkb kkbsc pfcvnq xcxvvx pbfnr gxsfth xchzh qxgtsm slgdh pqxvmx zdmk vkxsv khhf hdjqhg rrnt nrrll nbgklf rlrdk zjt rfpnz bqzmb gvnm mxxmpfs czpx ljl vcgtt vhf rdkx dvxq xhpnq hsxj chqbxcd njcqb fpcfdg shqqg mxzbbh vgz stdq htrzq dxhbtfmk fcrmdrf hhbf jslmtf dntjn pptnq tvfshn zkzls sjhds vgsgnkq fmj lxdjq pgxkln jtrgt mtdsv gsmbm bhkzzl tbthrm qdgnkl rkjjsl xktk lskhfb clvr jpljgr dgbz jpxs bsnnnhq lrbbmj nblsxmr fttbhdr sjkcx (contains eggs)\nltphvb vptd nnpz rdkx dvxq srk tvfshn xctm xchzh clvr vgfkcnr hdsm lxdjq gsmbm szpbd hnctrp bxvkx msmnzlm pptnq zvlpx hfxxf qlkhg svjsqv fttbhdr rldlxf jzqtjc lgjjh nchrg vlpzjs kmdgt kmhrsh dgbz fpcfdg nbzzpx vhgt nkgd kkbsc bvxkvt vvkrmf nrrll pfcvnq qlqrm fncgkl rrnt nfnfk vhf hqkxv kjxlsmb sjhds qjxxpr hxcshv gzlqd pqxvmx bthf qdgnkl htrzq sgxzx nlszr lskhfb (contains soy, peanuts, sesame)\nhdjqhg njcqb dgbz ltphvb nblsxmr kfgd nbzzpx sbkb nmgqt jqtlkmz pdkcrm rkjjsl jpljgr sjhds fttbhdr vlpzjs hfxxf jcfmh chxrpb sjkcx cbmbhft dgdq gzbhn kmdgt trq xfn zlzkh nzkr pjrlld nfnfk dntjn jslmtf kjxlsmb zkzls sgpncx bxvkx srk nlszr xchzh npmf mzvkx gsmbm lskhfb clvr hdsm vhgt ljl stq hhbf dlpnp qjxxpr mznbn bthf fpcfdg dvknd plqvl fncgkl lqvk mzj qdgnkl gvnm kgdqhg hnctrp tbthrm (contains dairy, fish, peanuts)\nzkzls nblsxmr bhkzzl ltzrz hqkxv pgxkln hxgl dxhbtfmk qjxxpr gcvqnz khtb gpscr dgpg nvtmrs stdq mlglz srk xphq xhpnq kkbsc zk bthf vlpzjs chqbxcd lzjhp vptd vzj nfnfk nbgklf xlrtg fttbhdr shqqg vhgt rldlxf mxzbbh bxvkx xdnbq xchzh hnctrp vgsgnkq rgtb zhzspb sjhds xbfht pdkcrm gjssscmd dzrzc mzj zjt hdsm tbthrm pqxvmx jtrgt jpljgr czpx gzbhn dvxq lqvk vmn lgjjh qqfnqb hdbjmn qdgnkl fflp jcfmh hdjqhg vvkrmf (contains dairy)\ngrgj zkzls xbfht dtp rjnmg zsbc nfnfk sgxzx jtrgt hdsm hhbf vmn kfgd tbthrm slgdh gsmbm kjxlsmb nkhrnpf xdnbq gzlqd qjxxpr vkxsv vzj mjrrfzz hfxxf ndpgv pfcvnq krd cbmbhft sjhds djlpcc clvr lrbbmj pptnq bqhjlt rgtb zhzspb kmhrsh vgz nzkr xchzh stdq sgpncx vlpzjs bthf jslmtf gcvqnz pgxkln hkhd mzvkx fvszcb gjssscmd lqvk shqqg dgdq qlqrm njcqb gtqqlgl qlkhg qxgtsm plqvl zk vhf fttbhdr khtb zlzkh rmblkb (contains peanuts, fish)\nclvr vcgtt lzfb kkbsc nfnfk lzjhp lgjjh trq bqzmb stdq xlrtg fxzj qqfnqb gvnm vkvg hxcshv kmr jzqtjc gtc ghrxm mmzg zdmk fflp krd vkxsv fncgkl jpxs dgbz mntmxzf fttbhdr zk lskhfb jcfmh dntjn hpvfng jqtlkmz qjxxpr zdd pqxvmx kgdqhg khhf kbrr mznbn jxbqcp zjt nnpz rgffc mxzbbh mttl qgcxhbm nbgklf sgpncx dvxq pbfnr xhpnq fcrmdrf szpbd dlpnp rldlxf nrrll shqqg pdkv sjhds khtb rfpnz xchzh nkgd xctm rdkx (contains dairy, nuts, eggs)\nzsbc nfnfk htrzq jpxs qxgtsm dgbz khhf hdjqhg fmj nlszr tbthrm qjxxpr dxhbtfmk rrnt gpscr mznbn nbgklf kncpfb dgdq fcrmdrf qlqrm jtrgt gtqqlgl clvr nchrg rgtb hdsm xchzh sgxzx fttbhdr zlzkh cszhgf xphq ltphvb pdkv cbmbhft djlpcc vgsgnkq xfn rldlxf msmnzlm vptd vhgt (contains nuts, wheat, sesame)\nnjcqb clvr pptnq vcgtt kmr xfn vlpzjs mznbn vhgt jxbqcp kfgd dvknd qlkhg nrrll fttbhdr hdsm bhkzzl svjsqv fcrmdrf xlrtg nfnfk hdjqhg stq nzkr msmnzlm rmblkb grgj mxxmpfs rlrdk khtb rfpnz dml fflp kbrr sjhds qjxxpr hfxxf jzqtjc gjssscmd mzj cszhgf szpbd dntjn nbgklf jslmtf qlqrm (contains soy, sesame)\nvkxsv zsbc mxzbbh nvtmrs plqvl pdkv mtdsv bsnnnhq hdsm pjrlld hxgl zdd kmdgt xdnbq kkbsc qxgtsm qdlczn qqfnqb vzj dntjn pdkcrm nfnfk fcrmdrf djlpcc hxcshv njcqb jslmtf lxdjq srk clvr cszhgf nblsxmr dgbz sbkb rldlxf lrbbmj lqvk ghrxm kbrr rgffc gzlqd tbthrm hsxj stdq sjhds nlszr nmgqt vmn vgfkcnr nzkr vkvg fttbhdr rrnt lzfb zvlpx xchzh fvszcb chxrpb fflp zk dzrjg gcvqnz jzqtjc hpvfng lskhfb rmblkb tvfshn pqxvmx vvkrmf nkhrnpf zlzkh svjsqv nbgklf fghm htrzq sgxzx vcgtt mmzg pptnq vlpzjs hfxxf dcqp (contains soy, nuts, sesame)\nfxzj stdq nmgqt hkhd sgxzx tvfshn clvr jzqtjc gcvqnz nvtmrs rmblkb kkbsc bnmttd nrrll slgdh cszhgf nkgd chqbxcd hpvfng pdkv lskhfb gtqqlgl hxgl xphq bqzmb bsnnnhq qjxxpr mxxmpfs zjt nblsxmr lzfb zdd jtrgt xcxvvx nbgklf mzvkx dtp nfnfk mxzbbh mttl hhbf gzbhn rdkx bhkzzl vgsgnkq vgfkcnr hdbjmn jslmtf krd sjhds vhgt qlqrm zdmk srk nnpz hqkxv zhzspb vkvg bvxkvt xchzh hsxj fttbhdr czpx ljl vhf rfpnz bqhjlt gvnm dcqp hnctrp mtdsv mjrrfzz qgcxhbm mznbn (contains fish)\nkbrt kncpfb vhf npmf kmr kbrr htrzq rldlxf xhpnq plqvl qgcxhbm mjrrfzz vvkrmf zdmk xctm gtc jxbqcp xktk hqkxv mlglz fpcfdg hkhd kgdqhg jzqtjc rfpnz shqqg grgj zk sgpncx gjssscmd qxgtsm czpx vhgt zkzls sgxzx xchzh lgjjh pjrlld djlpcc qdgnkl pgxkln fcrmdrf nlszr ltphvb ljl qlqrm xbfht hdsm bvxkvt clvr pqxvmx qjxxpr nkhrnpf fghm dlpnp vptd srk gsmbm zsbc hfxxf fttbhdr cbmbhft xfn xdnbq rdkx sjhds stq dgpg nbgklf (contains eggs, fish, sesame)\nxphq vgsgnkq kkbsc kjxlsmb stq qlqrm jpxs mmzg lxdjq xlrtg nchrg dxhbtfmk ljl ndpgv khhf rdkx qlkhg vmn lrbbmj gtc krd pfcvnq rgffc dntjn jslmtf dtp mttl lskhfb rfpnz jpljgr hnctrp fxzj rldlxf kfgd mznbn bshd hkhd gxsfth kbrt fflp lzjhp fttbhdr msmnzlm xchzh kncpfb sjhds vzj bqhjlt mxxmpfs svjsqv qjxxpr nbzzpx jxbqcp rrnt hdjqhg cbmbhft gvnm sbkb bhkzzl rmblkb qxgtsm stdq clvr gcvqnz hhbf hdsm nfnfk vgz bvxkvt rlrdk jzqtjc hpvfng nrrll zsbc mjrrfzz pdkcrm (contains nuts)\nvkxsv gzlqd zsbc htrzq qlqrm dlpnp jzqtjc hdsm ghrxm dvknd hfxxf vcgtt hkhd dcqp vgfkcnr zk rrnt svjsqv nblsxmr qjxxpr xcxvvx qdgnkl ltzrz lzjhp hxgl kkbsc mttl nfnfk ljl rgffc stq shqqg jtrgt nkhrnpf msmnzlm dtp mmzg hpvfng fmj nnpz vlpzjs vqd rlrdk mxzbbh mntmxzf rkjjsl cfhnhp zkzls sgxzx sjhds kbrt hqkxv jcfmh nmgqt fttbhdr nzkr qlkhg xctm szpbd stdq lqvk nlszr krd mjrrfzz clvr trq grgj hhbf vmn pptnq qxgtsm nbgklf dntjn nrrll lgjjh gtqqlgl dgpg xdnbq fcrmdrf djlpcc (contains nuts, eggs, peanuts)\n",
			"nfnfk,nbgklf,clvr,fttbhdr,qjxxpr,hdsm,sjhds,xchzh"},
	}

	for i, test := range tests {
		actual := part2(test.in)
		if actual != test.want {
			t.Errorf("#%d (part2): actual = %s, but want %s", i, actual, test.want)
		}
	}
}
