/*
 * Emby REST API
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyMediaModelEnumsVideoMediaTypes string

// List of Emby.Media.Model.Enums.VideoMediaTypes
const (
	UNKNOWN_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "Unknown"
	COPY_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "copy"
	FLV1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "flv1"
	H263_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "h263"
	H263P_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "h263p"
	H264_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "h264"
	HEVC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "hevc"
	MJPEG_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "mjpeg"
	MPEG1VIDEO_EmbyMediaModelEnumsVideoMediaTypes      EmbyMediaModelEnumsVideoMediaTypes = "mpeg1video"
	MPEG2VIDEO_EmbyMediaModelEnumsVideoMediaTypes      EmbyMediaModelEnumsVideoMediaTypes = "mpeg2video"
	MPEG4_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "mpeg4"
	MSVIDEO1_EmbyMediaModelEnumsVideoMediaTypes        EmbyMediaModelEnumsVideoMediaTypes = "msvideo1"
	THEORA_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "theora"
	VC1IMAGE_EmbyMediaModelEnumsVideoMediaTypes        EmbyMediaModelEnumsVideoMediaTypes = "vc1image"
	VC1_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "vc1"
	VP8_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "vp8"
	VP9_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "vp9"
	WMV1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "wmv1"
	WMV2_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "wmv2"
	WMV3_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "wmv3"
	_012V__EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "_012v"
	_4XM__EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "_4xm"
	_8BPS__EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "_8bps"
	A64_MULTI_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "a64_multi"
	A64_MULTI5_EmbyMediaModelEnumsVideoMediaTypes      EmbyMediaModelEnumsVideoMediaTypes = "a64_multi5"
	AASC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "aasc"
	AIC_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "aic"
	ALIAS_PIX_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "alias_pix"
	AMV_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "amv"
	ANM_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "anm"
	ANSI_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "ansi"
	APNG_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "apng"
	ASV1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "asv1"
	ASV2_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "asv2"
	AURA_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "aura"
	AURA2_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "aura2"
	AV1_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "av1"
	AVRN_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "avrn"
	AVRP_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "avrp"
	AVS_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "avs"
	AVUI_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "avui"
	AYUV_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "ayuv"
	BETHSOFTVID_EmbyMediaModelEnumsVideoMediaTypes     EmbyMediaModelEnumsVideoMediaTypes = "bethsoftvid"
	BFI_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "bfi"
	BINKVIDEO_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "binkvideo"
	BINTEXT_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "bintext"
	BITPACKED_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "bitpacked"
	BMP_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "bmp"
	BMV_VIDEO_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "bmv_video"
	BRENDER_PIX_EmbyMediaModelEnumsVideoMediaTypes     EmbyMediaModelEnumsVideoMediaTypes = "brender_pix"
	C93_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "c93"
	CAVS_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "cavs"
	CDGRAPHICS_EmbyMediaModelEnumsVideoMediaTypes      EmbyMediaModelEnumsVideoMediaTypes = "cdgraphics"
	CDXL_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "cdxl"
	CFHD_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "cfhd"
	CINEPAK_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "cinepak"
	CLEARVIDEO_EmbyMediaModelEnumsVideoMediaTypes      EmbyMediaModelEnumsVideoMediaTypes = "clearvideo"
	CLJR_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "cljr"
	CLLC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "cllc"
	CMV_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "cmv"
	CPIA_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "cpia"
	CSCD_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "cscd"
	CYUV_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "cyuv"
	DAALA_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "daala"
	DDS_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "dds"
	DFA_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "dfa"
	DIRAC_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "dirac"
	DNXHD_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "dnxhd"
	DPX_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "dpx"
	DSICINVIDEO_EmbyMediaModelEnumsVideoMediaTypes     EmbyMediaModelEnumsVideoMediaTypes = "dsicinvideo"
	DVVIDEO_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "dvvideo"
	DXA_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "dxa"
	DXTORY_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "dxtory"
	DXV_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "dxv"
	ESCAPE124_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "escape124"
	ESCAPE130_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "escape130"
	EXR_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "exr"
	FFV1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "ffv1"
	FFVHUFF_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "ffvhuff"
	FIC_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "fic"
	FITS_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "fits"
	FLASHSV_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "flashsv"
	FLASHSV2_EmbyMediaModelEnumsVideoMediaTypes        EmbyMediaModelEnumsVideoMediaTypes = "flashsv2"
	FLIC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "flic"
	FMVC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "fmvc"
	FRAPS_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "fraps"
	FRWU_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "frwu"
	G2M_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "g2m"
	GDV_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "gdv"
	GIF_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "gif"
	H261_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "h261"
	H263I_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "h263i"
	HAP_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "hap"
	HNM4VIDEO_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "hnm4video"
	HQ_HQA_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "hq_hqa"
	HQX_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "hqx"
	HUFFYUV_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "huffyuv"
	IDCIN_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "idcin"
	IDF_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "idf"
	IFF_ILBM_EmbyMediaModelEnumsVideoMediaTypes        EmbyMediaModelEnumsVideoMediaTypes = "iff_ilbm"
	INDEO2_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "indeo2"
	INDEO3_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "indeo3"
	INDEO4_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "indeo4"
	INDEO5_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "indeo5"
	INTERPLAYVIDEO_EmbyMediaModelEnumsVideoMediaTypes  EmbyMediaModelEnumsVideoMediaTypes = "interplayvideo"
	JPEG2000_EmbyMediaModelEnumsVideoMediaTypes        EmbyMediaModelEnumsVideoMediaTypes = "jpeg2000"
	JPEGLS_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "jpegls"
	JV_EmbyMediaModelEnumsVideoMediaTypes              EmbyMediaModelEnumsVideoMediaTypes = "jv"
	KGV1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "kgv1"
	KMVC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "kmvc"
	LAGARITH_EmbyMediaModelEnumsVideoMediaTypes        EmbyMediaModelEnumsVideoMediaTypes = "lagarith"
	LJPEG_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "ljpeg"
	LOCO_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "loco"
	M101_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "m101"
	MAD_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "mad"
	MAGICYUV_EmbyMediaModelEnumsVideoMediaTypes        EmbyMediaModelEnumsVideoMediaTypes = "magicyuv"
	MDEC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "mdec"
	MIMIC_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "mimic"
	MJPEGB_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "mjpegb"
	MMVIDEO_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "mmvideo"
	MOTIONPIXELS_EmbyMediaModelEnumsVideoMediaTypes    EmbyMediaModelEnumsVideoMediaTypes = "motionpixels"
	MSA1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "msa1"
	MSCC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "mscc"
	MSMPEG4V1_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "msmpeg4v1"
	MSMPEG4V2_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "msmpeg4v2"
	MSMPEG4V3_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "msmpeg4v3"
	MSRLE_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "msrle"
	MSS1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "mss1"
	MSS2_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "mss2"
	MSZH_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "mszh"
	MTS2_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "mts2"
	MVC1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "mvc1"
	MVC2_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "mvc2"
	MXPEG_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "mxpeg"
	NUV_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "nuv"
	PAF_VIDEO_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "paf_video"
	PAM_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "pam"
	PBM_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "pbm"
	PCX_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "pcx"
	PGM_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "pgm"
	PGMYUV_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "pgmyuv"
	PICTOR_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "pictor"
	PIXLET_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "pixlet"
	PNG_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "png"
	PPM_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "ppm"
	PRORES_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "prores"
	PSD_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "psd"
	PTX_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "ptx"
	QDRAW_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "qdraw"
	QPEG_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "qpeg"
	QTRLE_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "qtrle"
	R10K_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "r10k"
	R210_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "r210"
	RAWVIDEO_EmbyMediaModelEnumsVideoMediaTypes        EmbyMediaModelEnumsVideoMediaTypes = "rawvideo"
	RL2_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "rl2"
	ROQ_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "roq"
	RPZA_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "rpza"
	RSCC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "rscc"
	RV10_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "rv10"
	RV20_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "rv20"
	RV30_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "rv30"
	RV40_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "rv40"
	SANM_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "sanm"
	SCPR_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "scpr"
	SCREENPRESSO_EmbyMediaModelEnumsVideoMediaTypes    EmbyMediaModelEnumsVideoMediaTypes = "screenpresso"
	SGI_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "sgi"
	SGIRLE_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "sgirle"
	SHEERVIDEO_EmbyMediaModelEnumsVideoMediaTypes      EmbyMediaModelEnumsVideoMediaTypes = "sheervideo"
	SMACKVIDEO_EmbyMediaModelEnumsVideoMediaTypes      EmbyMediaModelEnumsVideoMediaTypes = "smackvideo"
	SMC_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "smc"
	SMVJPEG_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "smvjpeg"
	SNOW_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "snow"
	SP5X_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "sp5x"
	SPEEDHQ_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "speedhq"
	SRGC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "srgc"
	SUNRAST_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "sunrast"
	SVG_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "svg"
	SVQ1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "svq1"
	SVQ3_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "svq3"
	TARGA_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "targa"
	TARGA_Y216_EmbyMediaModelEnumsVideoMediaTypes      EmbyMediaModelEnumsVideoMediaTypes = "targa_y216"
	TDSC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "tdsc"
	TGQ_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "tgq"
	TGV_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "tgv"
	THP_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "thp"
	TIERTEXSEQVIDEO_EmbyMediaModelEnumsVideoMediaTypes EmbyMediaModelEnumsVideoMediaTypes = "tiertexseqvideo"
	TIFF_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "tiff"
	TMV_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "tmv"
	TQI_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "tqi"
	TRUEMOTION1_EmbyMediaModelEnumsVideoMediaTypes     EmbyMediaModelEnumsVideoMediaTypes = "truemotion1"
	TRUEMOTION2_EmbyMediaModelEnumsVideoMediaTypes     EmbyMediaModelEnumsVideoMediaTypes = "truemotion2"
	TRUEMOTION2RT_EmbyMediaModelEnumsVideoMediaTypes   EmbyMediaModelEnumsVideoMediaTypes = "truemotion2rt"
	TSCC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "tscc"
	TSCC2_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "tscc2"
	TXD_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "txd"
	ULTI_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "ulti"
	UTVIDEO_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "utvideo"
	V210_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "v210"
	V210X_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "v210x"
	V308_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "v308"
	V408_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "v408"
	V410_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "v410"
	VB_EmbyMediaModelEnumsVideoMediaTypes              EmbyMediaModelEnumsVideoMediaTypes = "vb"
	VBLE_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "vble"
	VCR1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "vcr1"
	VIXL_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "vixl"
	VMDVIDEO_EmbyMediaModelEnumsVideoMediaTypes        EmbyMediaModelEnumsVideoMediaTypes = "vmdvideo"
	VMNC_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "vmnc"
	VP3_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "vp3"
	VP5_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "vp5"
	VP6_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "vp6"
	VP6A_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "vp6a"
	VP6F_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "vp6f"
	VP7_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "vp7"
	WEBP_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "webp"
	WMV3IMAGE_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "wmv3image"
	WNV1_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "wnv1"
	WRAPPED_AVFRAME_EmbyMediaModelEnumsVideoMediaTypes EmbyMediaModelEnumsVideoMediaTypes = "wrapped_avframe"
	WS_VQA_EmbyMediaModelEnumsVideoMediaTypes          EmbyMediaModelEnumsVideoMediaTypes = "ws_vqa"
	XAN_WC3_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "xan_wc3"
	XAN_WC4_EmbyMediaModelEnumsVideoMediaTypes         EmbyMediaModelEnumsVideoMediaTypes = "xan_wc4"
	XBIN_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "xbin"
	XBM_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "xbm"
	XFACE_EmbyMediaModelEnumsVideoMediaTypes           EmbyMediaModelEnumsVideoMediaTypes = "xface"
	XPM_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "xpm"
	XWD_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "xwd"
	Y41P_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "y41p"
	YLC_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "ylc"
	YOP_EmbyMediaModelEnumsVideoMediaTypes             EmbyMediaModelEnumsVideoMediaTypes = "yop"
	YUV4_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "yuv4"
	ZEROCODEC_EmbyMediaModelEnumsVideoMediaTypes       EmbyMediaModelEnumsVideoMediaTypes = "zerocodec"
	ZLIB_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "zlib"
	ZMBV_EmbyMediaModelEnumsVideoMediaTypes            EmbyMediaModelEnumsVideoMediaTypes = "zmbv"
)
