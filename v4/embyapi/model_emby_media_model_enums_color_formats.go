/*
 * Emby Server REST API (BETA)
 *
 * Explore the Emby Server API
 *
 */
package embyapi

type EmbyMediaModelEnumsColorFormats string

// List of Emby.Media.Model.Enums.ColorFormats
const (
	UNKNOWN_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "Unknown"
	YUV420P_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "yuv420p"
	YUYV422_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "yuyv422"
	RGB24_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "rgb24"
	BGR24_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "bgr24"
	YUV422P_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "yuv422p"
	YUV444P_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "yuv444p"
	YUV410P_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "yuv410p"
	YUV411P_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "yuv411p"
	GRAY_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "gray"
	MONOW_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "monow"
	MONOB_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "monob"
	PAL8_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "pal8"
	YUVJ420P_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuvj420p"
	YUVJ422P_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuvj422p"
	YUVJ444P_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuvj444p"
	UYVY422_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "uyvy422"
	UYYVYY411_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "uyyvyy411"
	BGR8_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "bgr8"
	BGR4_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "bgr4"
	BGR4_BYTE_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "bgr4_byte"
	RGB8_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "rgb8"
	RGB4_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "rgb4"
	RGB4_BYTE_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "rgb4_byte"
	NV12_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "nv12"
	NV21_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "nv21"
	ARGB_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "argb"
	RGBA_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "rgba"
	ABGR_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "abgr"
	BGRA_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "bgra"
	GRAY16_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "gray16"
	YUV440P_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "yuv440p"
	YUVJ440P_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuvj440p"
	YUVA420P_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuva420p"
	RGB48_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "rgb48"
	RGB565_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "rgb565"
	RGB555_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "rgb555"
	BGR565_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "bgr565"
	BGR555_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "bgr555"
	VAAPI_MOCO_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "vaapi_moco"
	VAAPI_IDCT_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "vaapi_idct"
	VAAPI_VLD_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "vaapi_vld"
	YUV420P16_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv420p16"
	YUV422P16_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv422p16"
	YUV444P16_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv444p16"
	DXVA2_VLD_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "dxva2_vld"
	RGB444_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "rgb444"
	BGR444_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "bgr444"
	YA8_EmbyMediaModelEnumsColorFormats              EmbyMediaModelEnumsColorFormats = "ya8"
	BGR48_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "bgr48"
	YUV420P9_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuv420p9"
	YUV420P10_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv420p10"
	YUV422P10_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv422p10"
	YUV444P9_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuv444p9"
	YUV444P10_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv444p10"
	YUV422P9_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuv422p9"
	GBRP_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "gbrp"
	GBRP9_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "gbrp9"
	GBRP10_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "gbrp10"
	GBRP16_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "gbrp16"
	YUVA422P_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuva422p"
	YUVA444P_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuva444p"
	YUVA420P9_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuva420p9"
	YUVA422P9_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuva422p9"
	YUVA444P9_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuva444p9"
	YUVA420P10_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "yuva420p10"
	YUVA422P10_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "yuva422p10"
	YUVA444P10_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "yuva444p10"
	YUVA420P16_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "yuva420p16"
	YUVA422P16_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "yuva422p16"
	YUVA444P16_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "yuva444p16"
	VDPAU_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "vdpau"
	XYZ12_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "xyz12"
	NV16_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "nv16"
	NV20_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "nv20"
	RGBA64_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "rgba64"
	BGRA64_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "bgra64"
	YVYU422_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "yvyu422"
	YA16_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "ya16"
	GBRAP_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "gbrap"
	GBRAP16_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "gbrap16"
	QSV_EmbyMediaModelEnumsColorFormats              EmbyMediaModelEnumsColorFormats = "qsv"
	MMAL_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "mmal"
	D3D11VA_VLD_EmbyMediaModelEnumsColorFormats      EmbyMediaModelEnumsColorFormats = "d3d11va_vld"
	CUDA_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "cuda"
	_0RGB__EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "_0rgb"
	RGB0_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "rgb0"
	_0BGR__EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "_0bgr"
	BGR0_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "bgr0"
	YUV420P12_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv420p12"
	YUV420P14_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv420p14"
	YUV422P12_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv422p12"
	YUV422P14_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv422p14"
	YUV444P12_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv444p12"
	YUV444P14_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv444p14"
	GBRP12_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "gbrp12"
	GBRP14_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "gbrp14"
	YUVJ411P_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "yuvj411p"
	BAYER_BGGR8_EmbyMediaModelEnumsColorFormats      EmbyMediaModelEnumsColorFormats = "bayer_bggr8"
	BAYER_RGGB8_EmbyMediaModelEnumsColorFormats      EmbyMediaModelEnumsColorFormats = "bayer_rggb8"
	BAYER_GBRG8_EmbyMediaModelEnumsColorFormats      EmbyMediaModelEnumsColorFormats = "bayer_gbrg8"
	BAYER_GRBG8_EmbyMediaModelEnumsColorFormats      EmbyMediaModelEnumsColorFormats = "bayer_grbg8"
	BAYER_BGGR16_EmbyMediaModelEnumsColorFormats     EmbyMediaModelEnumsColorFormats = "bayer_bggr16"
	BAYER_RGGB16_EmbyMediaModelEnumsColorFormats     EmbyMediaModelEnumsColorFormats = "bayer_rggb16"
	BAYER_GBRG16_EmbyMediaModelEnumsColorFormats     EmbyMediaModelEnumsColorFormats = "bayer_gbrg16"
	BAYER_GRBG16_EmbyMediaModelEnumsColorFormats     EmbyMediaModelEnumsColorFormats = "bayer_grbg16"
	XVMC_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "xvmc"
	YUV440P10_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv440p10"
	YUV440P12_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "yuv440p12"
	AYUV64_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "ayuv64"
	VIDEOTOOLBOX_VLD_EmbyMediaModelEnumsColorFormats EmbyMediaModelEnumsColorFormats = "videotoolbox_vld"
	P010_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "p010"
	GBRAP12_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "gbrap12"
	GBRAP10_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "gbrap10"
	MEDIACODEC_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "mediacodec"
	GRAY12_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "gray12"
	GRAY10_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "gray10"
	GRAY14_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "gray14"
	P016_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "p016"
	D3D11_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "d3d11"
	GRAY9_EmbyMediaModelEnumsColorFormats            EmbyMediaModelEnumsColorFormats = "gray9"
	GBRPF32_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "gbrpf32"
	GBRAPF32_EmbyMediaModelEnumsColorFormats         EmbyMediaModelEnumsColorFormats = "gbrapf32"
	DRM_PRIME_EmbyMediaModelEnumsColorFormats        EmbyMediaModelEnumsColorFormats = "drm_prime"
	OPENCL_EmbyMediaModelEnumsColorFormats           EmbyMediaModelEnumsColorFormats = "opencl"
	GRAYF32_EmbyMediaModelEnumsColorFormats          EmbyMediaModelEnumsColorFormats = "grayf32"
	YUVA422P12_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "yuva422p12"
	YUVA444P12_EmbyMediaModelEnumsColorFormats       EmbyMediaModelEnumsColorFormats = "yuva444p12"
	NV24_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "nv24"
	NV42_EmbyMediaModelEnumsColorFormats             EmbyMediaModelEnumsColorFormats = "nv42"
)
