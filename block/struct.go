// NOTE: generated automatically by blendef for Blender v267.

package block

import (
	"fmt"
)

// BlenderVer is the version of Blender used when generating the files
// "parse.go" and "struct.go" of this package. Use blendef to regenerate these
// files if this version differs from the blend file's version.
const BlenderVer = 267

// Pointer is the memory address of a structure when it was written to disk.
type Pointer uint64

// Addr is a map from the memory address of a structure (when it was written to
// disk) to it's file block.
var Addr = make(map[uint64]*Block)

// Data translates the memory address into a usable pointer and returns it.
func (addr Pointer) Data() (data interface{}, err error) {
	blk, ok := Addr[uint64(addr)]
	if !ok {
		return nil, fmt.Errorf("Pointer.Data: unable to locate data for pointer %p.", addr)
	}
	return blk.Body, nil
}

// SDNA index: 0
type Link struct {
	Next Pointer // *Link
	Prev Pointer // *Link
}

// SDNA index: 1
type LinkData struct {
	Next Pointer // *LinkData
	Prev Pointer // *LinkData
	Data Pointer // *struct{}
}

// SDNA index: 2
type ListBase struct {
	First Pointer // *struct{}
	Last  Pointer // *struct{}
}

// SDNA index: 3
type Vec2s struct {
	X int16
	Y int16
}

// SDNA index: 4
type Vec2f struct {
	X float32
	Y float32
}

// SDNA index: 5
type Vec3f struct {
	X float32
	Y float32
	Z float32
}

// SDNA index: 6
type Rcti struct {
	Xmin int32
	Xmax int32
	Ymin int32
	Ymax int32
}

// SDNA index: 7
type Rctf struct {
	Xmin float32
	Xmax float32
	Ymin float32
	Ymax float32
}

// SDNA index: 8
type IDPropertyData struct {
	Pointer Pointer // *struct{}
	Group   ListBase
	Val     int32
	Val2    int32
}

// SDNA index: 9
type IDProperty struct {
	Next     Pointer // *IDProperty
	Prev     Pointer // *IDProperty
	Type     int8
	Subtype  int8
	Flag     int16
	Name     [64]int8
	Saved    int32
	Data     IDPropertyData
	Len      int32
	Totallen int32
}

// SDNA index: 10
type ID struct {
	Next       Pointer // *struct{}
	Prev       Pointer // *struct{}
	Newid      Pointer // *ID
	Lib        Pointer // *Library
	Name       [66]int8
	Pad        int16
	Us         int16
	Flag       int16
	Icon_id    int32
	Pad2       int32
	Properties Pointer // *IDProperty
}

// SDNA index: 11
type Library struct {
	Id         ID
	Idblock    Pointer // *ID
	Filedata   Pointer // *FileData
	Name       [1024]int8
	Filepath   [1024]int8
	Parent     Pointer // *Library
	Packedfile Pointer // *PackedFile
}

// SDNA index: 12
type PreviewImage struct {
	W                 [2]int32
	H                 [2]int32
	Changed           [2]int16
	Changed_timestamp [2]int16
	Rect              [2]Pointer // [2]*int32
	Gputexture        [2]Pointer // [2]*GPUTexture
}

// SDNA index: 13
type IpoDriver struct {
	Ob        Pointer // *Object
	Blocktype int16
	Adrcode   int16
	Type      int16
	Flag      int16
	Name      [128]int8
}

// SDNA index: 14
type IpoCurve struct {
	Next      Pointer // *IpoCurve
	Prev      Pointer // *IpoCurve
	Bp        Pointer // *BPoint
	Bezt      Pointer // *BezTriple
	Maxrct    Rctf
	Totrct    Rctf
	Blocktype int16
	Adrcode   int16
	Vartype   int16
	Totvert   int16
	Ipo       int16
	Extrap    int16
	Flag      int16
	Rt        int16
	Ymin      float32
	Ymax      float32
	Bitmask   int32
	Slide_min float32
	Slide_max float32
	Curval    float32
	Driver    Pointer // *IpoDriver
}

// SDNA index: 15
type Ipo struct {
	Id        ID
	Curve     ListBase
	Cur       Rctf
	Blocktype int16
	Showkey   int16
	Muteipo   int16
	Pad       int16
}

// SDNA index: 16
type KeyBlock struct {
	Next      Pointer // *KeyBlock
	Prev      Pointer // *KeyBlock
	Pos       float32
	Curval    float32
	Type      int16
	Pad1      int16
	Relative  int16
	Flag      int16
	Totelem   int32
	Uid       int32
	Data      Pointer // *struct{}
	Weights   Pointer // *float32
	Name      [64]int8
	Vgroup    [64]int8
	Slidermin float32
	Slidermax float32
}

// SDNA index: 17
type Key struct {
	Id       ID
	Adt      Pointer // *AnimData
	Refkey   Pointer // *KeyBlock
	Elemstr  [32]int8
	Elemsize int32
	Pad      int32
	Block    ListBase
	Ipo      Pointer // *Ipo
	From     Pointer // *ID
	Type     int16
	Totkey   int16
	Slurph   int16
	Flag     int16
	Ctime    float32
	Uidgen   int32
}

// SDNA index: 18
type TextLine struct {
	Next   Pointer // *TextLine
	Prev   Pointer // *TextLine
	Line   Pointer // *int8
	Format Pointer // *int8
	Len    int32
	Blen   int32
}

// SDNA index: 19
type Text struct {
	Id       ID
	Name     Pointer // *int8
	Flags    int32
	Nlines   int32
	Lines    ListBase
	Curl     Pointer // *TextLine
	Sell     Pointer // *TextLine
	Curc     int32
	Selc     int32
	Undo_buf Pointer // *int8
	Undo_pos int32
	Undo_len int32
	Compiled Pointer // *struct{}
	Mtime    float64
}

// SDNA index: 20
type PackedFile struct {
	Size int32
	Seek int32
	Data Pointer // *struct{}
}

// SDNA index: 21
type Camera struct {
	Id             ID
	Adt            Pointer // *AnimData
	Type           int8
	Dtx            int8
	Flag           int16
	Passepartalpha float32
	Clipsta        float32
	Clipend        float32
	Lens           float32
	Ortho_scale    float32
	Drawsize       float32
	Sensor_x       float32
	Sensor_y       float32
	Shiftx         float32
	Shifty         float32
	YF_dofdist     float32
	Ipo            Pointer // *Ipo
	Dof_ob         Pointer // *Object
	Sensor_fit     int8
	Pad            [7]int8
}

// SDNA index: 22
type ImageUser struct {
	Scene       Pointer // *Scene
	Framenr     int32
	Frames      int32
	Offset      int32
	Sfra        int32
	Fie_ima     int8
	Cycl        int8
	Ok          int8
	Pad         int8
	Multi_index int16
	Layer       int16
	Pass        int16
	Flag        int16
	Pad2        int32
}

// SDNA index: 23
type Image struct {
	Id                  ID
	Name                [1024]int8
	Ibufs               ListBase
	Gputexture          Pointer    // *GPUTexture
	Anim                Pointer    // *Anim
	Rr                  Pointer    // *RenderResult
	Renders             [8]Pointer // [8]*RenderResult
	Render_slot         int16
	Last_render_slot    int16
	Ok                  int16
	Flag                int16
	Source              int16
	Type                int16
	Lastframe           int32
	Tpageflag           int16
	Totbind             int16
	Xrep                int16
	Yrep                int16
	Twsta               int16
	Twend               int16
	Bindcode            int32
	Repbind             Pointer // *int32
	Packedfile          Pointer // *PackedFile
	Preview             Pointer // *PreviewImage
	Lastupdate          float32
	Lastused            int32
	Animspeed           int16
	Pad2                int16
	Gen_x               int32
	Gen_y               int32
	Gen_type            int8
	Gen_flag            int8
	Gen_pad             [2]int8
	Aspx                float32
	Aspy                float32
	Colorspace_settings ColorManagedColorspaceSettings
	Alpha_mode          int8
	Pad                 [7]int8
}

// SDNA index: 24
type MTex struct {
	Texco          int16
	Mapto          int16
	Maptoneg       int16
	Blendtype      int16
	Object         Pointer // *Object
	Tex            Pointer // *Tex
	Uvname         [64]int8
	Projx          int8
	Projy          int8
	Projz          int8
	Mapping        int8
	Ofs            [3]float32
	Size           [3]float32
	Rot            float32
	Texflag        int16
	Colormodel     int16
	Pmapto         int16
	Pmaptoneg      int16
	Normapspace    int16
	Which_output   int16
	Brush_map_mode int8
	Pad            [7]int8
	R              float32
	G              float32
	B              float32
	K              float32
	Def_var        float32
	Rt             float32
	Colfac         float32
	Varfac         float32
	Norfac         float32
	Dispfac        float32
	Warpfac        float32
	Colspecfac     float32
	Mirrfac        float32
	Alphafac       float32
	Difffac        float32
	Specfac        float32
	Emitfac        float32
	Hardfac        float32
	Raymirrfac     float32
	Translfac      float32
	Ambfac         float32
	Colemitfac     float32
	Colreflfac     float32
	Coltransfac    float32
	Densfac        float32
	Scatterfac     float32
	Reflfac        float32
	Timefac        float32
	Lengthfac      float32
	Clumpfac       float32
	Dampfac        float32
	Kinkfac        float32
	Roughfac       float32
	Padensfac      float32
	Gravityfac     float32
	Lifefac        float32
	Sizefac        float32
	Ivelfac        float32
	Fieldfac       float32
	Shadowfac      float32
	Zenupfac       float32
	Zendownfac     float32
	Blendfac       float32
}

// SDNA index: 25
type CBData struct {
	R   float32
	G   float32
	B   float32
	A   float32
	Pos float32
	Cur int32
}

// SDNA index: 26
type ColorBand struct {
	Flag    int16
	Tot     int16
	Cur     int16
	Ipotype int16
	Data    [32]CBData
}

// SDNA index: 27
type EnvMap struct {
	Object    Pointer    // *Object
	Ima       Pointer    // *Image
	Cube      [6]Pointer // [6]*ImBuf
	Imat      [4][4]float32
	Obimat    [3][3]float32
	Type      int16
	Stype     int16
	Clipsta   float32
	Clipend   float32
	Viewscale float32
	Notlay    int32
	Cuberes   int16
	Depth     int16
	Ok        int32
	Lastframe int32
	Recalc    int16
	Lastsize  int16
}

// SDNA index: 28
type PointDensity struct {
	Flag                int16
	Falloff_type        int16
	Falloff_softness    float32
	Radius              float32
	Source              int16
	Color_source        int16
	Totpoints           int32
	Pdpad               int32
	Object              Pointer // *Object
	Psys                int32
	Psys_cache_space    int16
	Ob_cache_space      int16
	Point_tree          Pointer // *struct{}
	Point_data          Pointer // *float32
	Noise_size          float32
	Noise_depth         int16
	Noise_influence     int16
	Noise_basis         int16
	Pdpad3              [3]int16
	Noise_fac           float32
	Speed_scale         float32
	Falloff_speed_scale float32
	Pdpad2              float32
	Coba                Pointer // *ColorBand
	Falloff_curve       Pointer // *CurveMapping
}

// SDNA index: 29
type VoxelData struct {
	Resol          [3]int32
	Interp_type    int32
	File_format    int16
	Flag           int16
	Extend         int16
	Smoked_type    int16
	Data_type      int16
	Pad            int16
	_pad           int32
	Object         Pointer // *Object
	Int_multiplier float32
	Still_frame    int32
	Source_path    [1024]int8
	Dataset        Pointer // *float32
	Cachedframe    int32
	Ok             int32
}

// SDNA index: 30
type OceanTex struct {
	Object   Pointer // *Object
	Oceanmod [64]int8
	Output   int32
	Pad      int32
}

// SDNA index: 31
type Tex struct {
	Id            ID
	Adt           Pointer // *AnimData
	Noisesize     float32
	Turbul        float32
	Bright        float32
	Contrast      float32
	Saturation    float32
	Rfac          float32
	Gfac          float32
	Bfac          float32
	Filtersize    float32
	Pad2          float32
	Mg_H          float32
	Mg_lacunarity float32
	Mg_octaves    float32
	Mg_offset     float32
	Mg_gain       float32
	Dist_amount   float32
	Ns_outscale   float32
	Vn_w1         float32
	Vn_w2         float32
	Vn_w3         float32
	Vn_w4         float32
	Vn_mexp       float32
	Vn_distm      int16
	Vn_coltype    int16
	Noisedepth    int16
	Noisetype     int16
	Noisebasis    int16
	Noisebasis2   int16
	Imaflag       int16
	Flag          int16
	Type          int16
	Stype         int16
	Cropxmin      float32
	Cropymin      float32
	Cropxmax      float32
	Cropymax      float32
	Texfilter     int32
	Afmax         int32
	Xrepeat       int16
	Yrepeat       int16
	Extend        int16
	Fie_ima       int16
	Len           int32
	Frames        int32
	Offset        int32
	Sfra          int32
	Checkerdist   float32
	Nabla         float32
	Pad1          float32
	Iuser         ImageUser
	Nodetree      Pointer // *BNodeTree
	Ipo           Pointer // *Ipo
	Ima           Pointer // *Image
	Coba          Pointer // *ColorBand
	Env           Pointer // *EnvMap
	Preview       Pointer // *PreviewImage
	Pd            Pointer // *PointDensity
	Vd            Pointer // *VoxelData
	Ot            Pointer // *OceanTex
	Use_nodes     int8
	Pad           [7]int8
}

// SDNA index: 32
type TexMapping struct {
	Loc     [3]float32
	Rot     [3]float32
	Size    [3]float32
	Flag    int32
	Projx   int8
	Projy   int8
	Projz   int8
	Mapping int8
	Pad     int32
	Mat     [4][4]float32
	Min     [3]float32
	Max     [3]float32
	Ob      Pointer // *Object
}

// SDNA index: 33
type ColorMapping struct {
	Coba         ColorBand
	Bright       float32
	Contrast     float32
	Saturation   float32
	Flag         int32
	Blend_color  [3]float32
	Blend_factor float32
	Blend_type   int32
	Pad          [3]int32
}

// SDNA index: 34
type Lamp struct {
	Id                      ID
	Adt                     Pointer // *AnimData
	Type                    int16
	Flag                    int16
	Mode                    int32
	Colormodel              int16
	Totex                   int16
	R                       float32
	G                       float32
	B                       float32
	K                       float32
	Shdwr                   float32
	Shdwg                   float32
	Shdwb                   float32
	Shdwpad                 float32
	Energy                  float32
	Dist                    float32
	Spotsize                float32
	Spotblend               float32
	Haint                   float32
	Att1                    float32
	Att2                    float32
	Curfalloff              Pointer // *CurveMapping
	Falloff_type            int16
	Pad2                    int16
	Clipsta                 float32
	Clipend                 float32
	Shadspotsize            float32
	Bias                    float32
	Soft                    float32
	Compressthresh          float32
	Bleedbias               float32
	Pad5                    [2]float32
	Bufsize                 int16
	Samp                    int16
	Buffers                 int16
	Filtertype              int16
	Bufflag                 int8
	Buftype                 int8
	Ray_samp                int16
	Ray_sampy               int16
	Ray_sampz               int16
	Ray_samp_type           int16
	Area_shape              int16
	Area_size               float32
	Area_sizey              float32
	Area_sizez              float32
	Adapt_thresh            float32
	Ray_samp_method         int16
	Shadowmap_type          int16
	Texact                  int16
	Shadhalostep            int16
	Sun_effect_type         int16
	Skyblendtype            int16
	Horizon_brightness      float32
	Spread                  float32
	Sun_brightness          float32
	Sun_size                float32
	Backscattered_light     float32
	Sun_intensity           float32
	Atm_turbidity           float32
	Atm_inscattering_factor float32
	Atm_extinction_factor   float32
	Atm_distance_factor     float32
	Skyblendfac             float32
	Sky_exposure            float32
	Shadow_frustum_size     float32
	Sky_colorspace          int16
	Pad4                    [2]int8
	Ipo                     Pointer     // *Ipo
	Mtex                    [18]Pointer // [18]*MTex
	Pr_texture              int16
	Use_nodes               int16
	Pad6                    [4]int8
	Preview                 Pointer // *PreviewImage
	Nodetree                Pointer // *BNodeTree
}

// SDNA index: 35
type VolumeSettings struct {
	Density             float32
	Emission            float32
	Scattering          float32
	Reflection          float32
	Emission_col        [3]float32
	Transmission_col    [3]float32
	Reflection_col      [3]float32
	Density_scale       float32
	Depth_cutoff        float32
	Asymmetry           float32
	Stepsize_type       int16
	Shadeflag           int16
	Shade_type          int16
	Precache_resolution int16
	Stepsize            float32
	Ms_diff             float32
	Ms_intensity        float32
	Ms_spread           float32
}

// SDNA index: 36
type GameSettings struct {
	Flag             int32
	Alpha_blend      int32
	Face_orientation int32
	Pad1             int32
}

// SDNA index: 37
type Material struct {
	Id               ID
	Adt              Pointer // *AnimData
	Material_type    int16
	Flag             int16
	R                float32
	G                float32
	B                float32
	Specr            float32
	Specg            float32
	Specb            float32
	Mirr             float32
	Mirg             float32
	Mirb             float32
	Ambr             float32
	Ambb             float32
	Ambg             float32
	Amb              float32
	Emit             float32
	Ang              float32
	Spectra          float32
	Ray_mirror       float32
	Alpha            float32
	Ref              float32
	Spec             float32
	Zoffs            float32
	Add              float32
	Translucency     float32
	Vol              VolumeSettings
	Game             GameSettings
	Fresnel_mir      float32
	Fresnel_mir_i    float32
	Fresnel_tra      float32
	Fresnel_tra_i    float32
	Filter           float32
	Tx_limit         float32
	Tx_falloff       float32
	Ray_depth        int16
	Ray_depth_tra    int16
	Har              int16
	Seed1            int8
	Seed2            int8
	Gloss_mir        float32
	Gloss_tra        float32
	Samp_gloss_mir   int16
	Samp_gloss_tra   int16
	Adapt_thresh_mir float32
	Adapt_thresh_tra float32
	Aniso_gloss_mir  float32
	Dist_mir         float32
	Fadeto_mir       int16
	Shade_flag       int16
	Mode             int32
	Mode_l           int32
	Flarec           int16
	Starc            int16
	Linec            int16
	Ringc            int16
	Hasize           float32
	Flaresize        float32
	Subsize          float32
	Flareboost       float32
	Strand_sta       float32
	Strand_end       float32
	Strand_ease      float32
	Strand_surfnor   float32
	Strand_min       float32
	Strand_widthfade float32
	Strand_uvname    [64]int8
	Sbias            float32
	Lbias            float32
	Shad_alpha       float32
	Septex           int32
	Rgbsel           int8
	Texact           int8
	Pr_type          int8
	Use_nodes        int8
	Pr_lamp          int16
	Pr_texture       int16
	Ml_flag          int16
	Mapflag          int8
	Pad              int8
	Diff_shader      int16
	Spec_shader      int16
	Roughness        float32
	Refrac           float32
	Param            [4]float32
	Rms              float32
	Darkness         float32
	Texco            int16
	Mapto            int16
	Ramp_col         Pointer // *ColorBand
	Ramp_spec        Pointer // *ColorBand
	Rampin_col       int8
	Rampin_spec      int8
	Rampblend_col    int8
	Rampblend_spec   int8
	Ramp_show        int16
	Pad3             int16
	Rampfac_col      float32
	Rampfac_spec     float32
	Mtex             [18]Pointer // [18]*MTex
	Nodetree         Pointer     // *BNodeTree
	Ipo              Pointer     // *Ipo
	Group            Pointer     // *Group
	Preview          Pointer     // *PreviewImage
	Friction         float32
	Fh               float32
	Reflect          float32
	Fhdist           float32
	Xyfrict          float32
	Dynamode         int16
	Pad2             int16
	Sss_radius       [3]float32
	Sss_col          [3]float32
	Sss_error        float32
	Sss_scale        float32
	Sss_ior          float32
	Sss_colfac       float32
	Sss_texfac       float32
	Sss_front        float32
	Sss_back         float32
	Sss_flag         int16
	Sss_preset       int16
	Mapto_textured   int32
	Shadowonly_flag  int16
	Index            int16
	Vcol_alpha       int16
	Pad4             [3]int16
	Gpumaterial      ListBase
}

// SDNA index: 38
type VFont struct {
	Id         ID
	Name       [1024]int8
	Data       Pointer // *VFontData
	Packedfile Pointer // *PackedFile
	Temp_pf    Pointer // *PackedFile
}

// SDNA index: 39
type MetaElem struct {
	Next    Pointer // *MetaElem
	Prev    Pointer // *MetaElem
	Bb      Pointer // *BoundBox
	Type    int16
	Flag    int16
	Selcol1 int16
	Selcol2 int16
	X       float32
	Y       float32
	Z       float32
	Quat    [4]float32
	Expx    float32
	Expy    float32
	Expz    float32
	Rad     float32
	Rad2    float32
	S       float32
	Len     float32
	Mat     Pointer // *float32
	Imat    Pointer // *float32
}

// SDNA index: 40
type MetaBall struct {
	Id         ID
	Adt        Pointer // *AnimData
	Bb         Pointer // *BoundBox
	Elems      ListBase
	Disp       ListBase
	Editelems  Pointer // *ListBase
	Ipo        Pointer // *Ipo
	Mat        Pointer // **Material
	Flag       int8
	Flag2      int8
	Totcol     int16
	Texflag    int16
	Pad        int16
	Loc        [3]float32
	Size       [3]float32
	Rot        [3]float32
	Wiresize   float32
	Rendersize float32
	Thresh     float32
	Lastelem   Pointer // *MetaElem
}

// SDNA index: 41
type BezTriple struct {
	Vec    [3][3]float32
	Alfa   float32
	Weight float32
	Radius float32
	Ipo    int16
	H1     int8
	H2     int8
	F1     int8
	F2     int8
	F3     int8
	Hide   int8
}

// SDNA index: 42
type BPoint struct {
	Vec    [4]float32
	Alfa   float32
	Weight float32
	F1     int16
	Hide   int16
	Radius float32
	Pad    float32
}

// SDNA index: 43
type Nurb struct {
	Next          Pointer // *Nurb
	Prev          Pointer // *Nurb
	Type          int16
	Mat_nr        int16
	Hide          int16
	Flag          int16
	Pntsu         int16
	Pntsv         int16
	Resolu        int16
	Resolv        int16
	Orderu        int16
	Orderv        int16
	Flagu         int16
	Flagv         int16
	Knotsu        Pointer // *float32
	Knotsv        Pointer // *float32
	Bp            Pointer // *BPoint
	Bezt          Pointer // *BezTriple
	Tilt_interp   int16
	Radius_interp int16
	Charidx       int32
}

// SDNA index: 44
type CharInfo struct {
	Kern   int16
	Mat_nr int16
	Flag   int8
	Pad    int8
	Pad2   int16
}

// SDNA index: 45
type TextBox struct {
	X float32
	Y float32
	W float32
	H float32
}

// SDNA index: 46
type EditNurb struct {
	Nurbs    ListBase
	Keyindex Pointer // *GHash
	Shapenr  int32
	Pad      [4]int8
}

// SDNA index: 47
type Curve struct {
	Id              ID
	Adt             Pointer // *AnimData
	Bb              Pointer // *BoundBox
	Nurb            ListBase
	Disp            ListBase
	Editnurb        Pointer // *EditNurb
	Bevobj          Pointer // *Object
	Taperobj        Pointer // *Object
	Textoncurve     Pointer // *Object
	Ipo             Pointer // *Ipo
	Path            Pointer // *Path
	Key             Pointer // *Key
	Mat             Pointer // **Material
	Bev             ListBase
	Loc             [3]float32
	Size            [3]float32
	Rot             [3]float32
	Type            int16
	Texflag         int16
	Drawflag        int16
	Twist_mode      int16
	Twist_smooth    float32
	Smallcaps_scale float32
	Pathlen         int32
	Bevresol        int16
	Totcol          int16
	Flag            int32
	Width           float32
	Ext1            float32
	Ext2            float32
	Resolu          int16
	Resolv          int16
	Resolu_ren      int16
	Resolv_ren      int16
	Actnu           int32
	Lastsel         Pointer // *struct{}
	Len             int16
	Lines           int16
	Pos             int16
	Spacemode       int16
	Spacing         float32
	Linedist        float32
	Shear           float32
	Fsize           float32
	Wordspace       float32
	Ulpos           float32
	Ulheight        float32
	Xof             float32
	Yof             float32
	Linewidth       float32
	Str             Pointer // *int8
	Selboxes        Pointer // *SelBox
	Editfont        Pointer // *EditFont
	Family          [24]int8
	Vfont           Pointer // *VFont
	Vfontb          Pointer // *VFont
	Vfonti          Pointer // *VFont
	Vfontbi         Pointer // *VFont
	Sepchar         int32
	Ctime           float32
	Totbox          int32
	Actbox          int32
	Tb              Pointer // *TextBox
	Selstart        int32
	Selend          int32
	Strinfo         Pointer // *CharInfo
	Curinfo         CharInfo
	Bevfac1         float32
	Bevfac2         float32
}

// SDNA index: 48
type Mesh struct {
	Id          ID
	Adt         Pointer // *AnimData
	Bb          Pointer // *BoundBox
	Ipo         Pointer // *Ipo
	Key         Pointer // *Key
	Mat         Pointer // **Material
	Mselect     Pointer // *MSelect
	Mpoly       Pointer // *MPoly
	Mtpoly      Pointer // *MTexPoly
	Mloop       Pointer // *MLoop
	Mloopuv     Pointer // *MLoopUV
	Mloopcol    Pointer // *MLoopCol
	Mface       Pointer // *MFace
	Mtface      Pointer // *MTFace
	Tface       Pointer // *TFace
	Mvert       Pointer // *MVert
	Medge       Pointer // *MEdge
	Dvert       Pointer // *MDeformVert
	Mcol        Pointer // *MCol
	Texcomesh   Pointer // *Mesh
	Edit_btmesh Pointer // *BMEditMesh
	Vdata       CustomData
	Edata       CustomData
	Fdata       CustomData
	Pdata       CustomData
	Ldata       CustomData
	Totvert     int32
	Totedge     int32
	Totface     int32
	Totselect   int32
	Totpoly     int32
	Totloop     int32
	Act_face    int32
	Loc         [3]float32
	Size        [3]float32
	Rot         [3]float32
	Drawflag    int32
	Texflag     int16
	Pad2        [3]int16
	Smoothresh  int16
	Flag        int16
	Cd_flag     int8
	Pad         int8
	Subdiv      int8
	Subdivr     int8
	Subsurftype int8
	Editflag    int8
	Totcol      int16
	Mr          Pointer // *Multires
}

// SDNA index: 49
type TFace struct {
	Tpage  Pointer // *struct{}
	Uv     [4][2]float32
	Col    [4]int32
	Flag   int8
	Transp int8
	Mode   int16
	Tile   int16
	Unwrap int16
}

// SDNA index: 50
type MFace struct {
	V1     int32
	V2     int32
	V3     int32
	V4     int32
	Mat_nr int16
	Edcode int8
	Flag   int8
}

// SDNA index: 51
type MEdge struct {
	V1      int32
	V2      int32
	Crease  int8
	Bweight int8
	Flag    int16
}

// SDNA index: 52
type MDeformWeight struct {
	Def_nr int32
	Weight float32
}

// SDNA index: 53
type MDeformVert struct {
	Dw        Pointer // *MDeformWeight
	Totweight int32
	Flag      int32
}

// SDNA index: 54
type MVert struct {
	Co      [3]float32
	No      [3]int16
	Flag    int8
	Bweight int8
}

// SDNA index: 55
type MCol struct {
	A int8
	R int8
	G int8
	B int8
}

// SDNA index: 56
type MPoly struct {
	Loopstart int32
	Totloop   int32
	Mat_nr    int16
	Flag      int8
	Pad       int8
}

// SDNA index: 57
type MLoop struct {
	V int32
	E int32
}

// SDNA index: 58
type MTexPoly struct {
	Tpage  Pointer // *Image
	Flag   int8
	Transp int8
	Mode   int16
	Tile   int16
	Pad    int16
}

// SDNA index: 59
type MLoopUV struct {
	Uv   [2]float32
	Flag int32
}

// SDNA index: 60
type MLoopCol struct {
	R int8
	G int8
	B int8
	A int8
}

// SDNA index: 61
type MSelect struct {
	Index int32
	Type  int32
}

// SDNA index: 62
type MTFace struct {
	Uv     [4][2]float32
	Tpage  Pointer // *Image
	Flag   int8
	Transp int8
	Mode   int16
	Tile   int16
	Unwrap int16
}

// SDNA index: 63
type MFloatProperty struct {
	F float32
}

// SDNA index: 64
type MIntProperty struct {
	I int32
}

// SDNA index: 65
type MStringProperty struct {
	S     [255]int8
	S_len int8
}

// SDNA index: 66
type OrigSpaceFace struct {
	Uv [4][2]float32
}

// SDNA index: 67
type OrigSpaceLoop struct {
	Uv [2]float32
}

// SDNA index: 68
type MDisps struct {
	Totdisp int32
	Level   int32
	Disps   func() float32
	Hidden  Pointer // *int32
}

// SDNA index: 69
type MultiresCol struct {
	A float32
	R float32
	G float32
	B float32
}

// SDNA index: 70
type MultiresColFace struct {
	Col [4]MultiresCol
}

// SDNA index: 71
type MultiresFace struct {
	V      [4]int32
	Mid    int32
	Flag   int8
	Mat_nr int8
	Pad    [2]int8
}

// SDNA index: 72
type MultiresEdge struct {
	V   [2]int32
	Mid int32
}

// SDNA index: 73
type MultiresLevel struct {
	Next     Pointer // *MultiresLevel
	Prev     Pointer // *MultiresLevel
	Faces    Pointer // *MultiresFace
	Colfaces Pointer // *MultiresColFace
	Edges    Pointer // *MultiresEdge
	Totvert  int32
	Totface  int32
	Totedge  int32
	Pad      int32
	Verts    Pointer // *MVert
}

// SDNA index: 74
type Multires struct {
	Levels       ListBase
	Verts        Pointer // *MVert
	Level_count  int8
	Current      int8
	Newlvl       int8
	Edgelvl      int8
	Pinlvl       int8
	Renderlvl    int8
	Use_col      int8
	Flag         int8
	Vdata        CustomData
	Fdata        CustomData
	Edge_flags   Pointer // *int16
	Edge_creases Pointer // *int8
}

// SDNA index: 75
type MRecast struct {
	I int32
}

// SDNA index: 76
type GridPaintMask struct {
	Data  Pointer // *float32
	Level int32
	Pad   int32
}

// SDNA index: 77
type MVertSkin struct {
	Radius [3]float32
	Flag   int32
}

// SDNA index: 78
type FreestyleEdge struct {
	Flag int8
	Pad  [3]int8
}

// SDNA index: 79
type FreestyleFace struct {
	Flag int8
	Pad  [3]int8
}

// SDNA index: 80
type ModifierData struct {
	Next       Pointer // *ModifierData
	Prev       Pointer // *ModifierData
	Type       int32
	Mode       int32
	Stackindex int32
	Pad        int32
	Name       [64]int8
	Scene      Pointer // *Scene
	Error      Pointer // *int8
}

// SDNA index: 81
type MappingInfoModifierData struct {
	Modifier     ModifierData
	Texture      Pointer // *Tex
	Map_object   Pointer // *Object
	Uvlayer_name [64]int8
	Uvlayer_tmp  int32
	Texmapping   int32
}

// SDNA index: 82
type SubsurfModifierData struct {
	Modifier     ModifierData
	SubdivType   int16
	Levels       int16
	RenderLevels int16
	Flags        int16
	EmCache      Pointer // *struct{}
	MCache       Pointer // *struct{}
}

// SDNA index: 83
type LatticeModifierData struct {
	Modifier ModifierData
	Object   Pointer // *Object
	Name     [64]int8
	Strength float32
	Pad      [4]int8
}

// SDNA index: 84
type CurveModifierData struct {
	Modifier ModifierData
	Object   Pointer // *Object
	Name     [64]int8
	Defaxis  int16
	Pad      [6]int8
}

// SDNA index: 85
type BuildModifierData struct {
	Modifier  ModifierData
	Start     float32
	Length    float32
	Randomize int32
	Seed      int32
}

// SDNA index: 86
type MaskModifierData struct {
	Modifier ModifierData
	Ob_arm   Pointer // *Object
	Vgroup   [64]int8
	Mode     int32
	Flag     int32
}

// SDNA index: 87
type ArrayModifierData struct {
	Modifier    ModifierData
	Start_cap   Pointer // *Object
	End_cap     Pointer // *Object
	Curve_ob    Pointer // *Object
	Offset_ob   Pointer // *Object
	Offset      [3]float32
	Scale       [3]float32
	Length      float32
	Merge_dist  float32
	Fit_type    int32
	Offset_type int32
	Flags       int32
	Count       int32
}

// SDNA index: 88
type MirrorModifierData struct {
	Modifier  ModifierData
	Axis      int16
	Flag      int16
	Tolerance float32
	Mirror_ob Pointer // *Object
}

// SDNA index: 89
type EdgeSplitModifierData struct {
	Modifier    ModifierData
	Split_angle float32
	Flags       int32
}

// SDNA index: 90
type BevelModifierData struct {
	Modifier    ModifierData
	Value       float32
	Res         int32
	Pad         int32
	Flags       int16
	Val_flags   int16
	Lim_flags   int16
	E_flags     int16
	Bevel_angle float32
	Defgrp_name [64]int8
}

// SDNA index: 91
type BMeshModifierData struct {
	Modifier ModifierData
	Pad      float32
	Type     int32
}

// SDNA index: 92
type SmokeModifierData struct {
	Modifier ModifierData
	Domain   Pointer // *SmokeDomainSettings
	Flow     Pointer // *SmokeFlowSettings
	Coll     Pointer // *SmokeCollSettings
	Time     float32
	Type     int32
}

// SDNA index: 93
type DisplaceModifierData struct {
	Modifier     ModifierData
	Texture      Pointer // *Tex
	Map_object   Pointer // *Object
	Uvlayer_name [64]int8
	Uvlayer_tmp  int32
	Texmapping   int32
	Strength     float32
	Direction    int32
	Defgrp_name  [64]int8
	Midlevel     float32
	Pad          int32
}

// SDNA index: 94
type UVProjectModifierData struct {
	Modifier       ModifierData
	Projectors     [10]Pointer // [10]*Object
	Image          Pointer     // *Image
	Flags          int32
	Num_projectors int32
	Aspectx        float32
	Aspecty        float32
	Scalex         float32
	Scaley         float32
	Uvlayer_name   [64]int8
	Uvlayer_tmp    int32
	Pad            int32
}

// SDNA index: 95
type DecimateModifierData struct {
	Modifier    ModifierData
	Percent     float32
	Iter        int16
	Pad         int16
	Angle       float32
	Defgrp_name [64]int8
	Flag        int16
	Mode        int16
	Face_count  int32
	Pad2        int32
}

// SDNA index: 96
type SmoothModifierData struct {
	Modifier    ModifierData
	Fac         float32
	Defgrp_name [64]int8
	Flag        int16
	Repeat      int16
}

// SDNA index: 97
type CastModifierData struct {
	Modifier    ModifierData
	Object      Pointer // *Object
	Fac         float32
	Radius      float32
	Size        float32
	Defgrp_name [64]int8
	Flag        int16
	Type        int16
}

// SDNA index: 98
type WaveModifierData struct {
	Modifier     ModifierData
	Texture      Pointer // *Tex
	Map_object   Pointer // *Object
	Uvlayer_name [64]int8
	Uvlayer_tmp  int32
	Texmapping   int32
	Objectcenter Pointer // *Object
	Defgrp_name  [64]int8
	Flag         int16
	Pad          int16
	Startx       float32
	Starty       float32
	Height       float32
	Width        float32
	Narrow       float32
	Speed        float32
	Damp         float32
	Falloff      float32
	Timeoffs     float32
	Lifetime     float32
	Pad1         float32
}

// SDNA index: 99
type ArmatureModifierData struct {
	Modifier    ModifierData
	Deformflag  int16
	Multi       int16
	Pad2        int32
	Object      Pointer // *Object
	PrevCos     Pointer // *float32
	Defgrp_name [64]int8
}

// SDNA index: 100
type HookModifierData struct {
	Modifier  ModifierData
	Object    Pointer // *Object
	Subtarget [64]int8
	Parentinv [4][4]float32
	Cent      [3]float32
	Falloff   float32
	Indexar   Pointer // *int32
	Totindex  int32
	Force     float32
	Name      [64]int8
}

// SDNA index: 101
type SoftbodyModifierData struct {
	Modifier ModifierData
}

// SDNA index: 102
type ClothModifierData struct {
	Modifier    ModifierData
	Scene       Pointer // *Scene
	ClothObject Pointer // *Cloth
	Sim_parms   Pointer // *ClothSimSettings
	Coll_parms  Pointer // *ClothCollSettings
	Point_cache Pointer // *PointCache
	Ptcaches    ListBase
}

// SDNA index: 103
type CollisionModifierData struct {
	Modifier     ModifierData
	X            Pointer // *MVert
	Xnew         Pointer // *MVert
	Xold         Pointer // *MVert
	Current_xnew Pointer // *MVert
	Current_x    Pointer // *MVert
	Current_v    Pointer // *MVert
	Mfaces       Pointer // *MFace
	Numverts     int32
	Numfaces     int32
	Time_x       float32
	Time_xnew    float32
	Bvhtree      Pointer // *BVHTree
}

// SDNA index: 104
type SurfaceModifierData struct {
	Modifier ModifierData
	X        Pointer // *MVert
	V        Pointer // *MVert
	Dm       Pointer // *DerivedMesh
	Bvhtree  Pointer // *BVHTreeFromMesh
	Cfra     int32
	Numverts int32
}

// SDNA index: 105
type BooleanModifierData struct {
	Modifier  ModifierData
	Object    Pointer // *Object
	Operation int32
	Pad       int32
}

// SDNA index: 106
type MDefInfluence struct {
	Vertex int32
	Weight float32
}

// SDNA index: 107
type MDefCell struct {
	Offset       int32
	Totinfluence int32
}

// SDNA index: 108
type MeshDeformModifierData struct {
	Modifier       ModifierData
	Object         Pointer // *Object
	Defgrp_name    [64]int8
	Gridsize       int16
	Flag           int16
	Mode           int16
	Pad            int16
	Bindinfluences Pointer // *MDefInfluence
	Bindoffsets    Pointer // *int32
	Bindcagecos    Pointer // *float32
	Totvert        int32
	Totcagevert    int32
	Dyngrid        Pointer // *MDefCell
	Dyninfluences  Pointer // *MDefInfluence
	Dynverts       Pointer // *int32
	Pad2           Pointer // *int32
	Dyngridsize    int32
	Totinfluence   int32
	Dyncellmin     [3]float32
	Dyncellwidth   float32
	Bindmat        [4][4]float32
	Bindweights    Pointer // *float32
	Bindcos        Pointer // *float32
	Bindfunc       func()
}

// SDNA index: 109
type ParticleSystemModifierData struct {
	Modifier  ModifierData
	Psys      Pointer // *ParticleSystem
	Dm        Pointer // *DerivedMesh
	Totdmvert int32
	Totdmedge int32
	Totdmface int32
	Flag      int16
	Rt        int16
}

// SDNA index: 110
type ParticleInstanceModifierData struct {
	Modifier        ModifierData
	Ob              Pointer // *Object
	Psys            int16
	Flag            int16
	Axis            int16
	Rt              int16
	Position        float32
	Random_position float32
}

// SDNA index: 111
type ExplodeModifierData struct {
	Modifier ModifierData
	Facepa   Pointer // *int32
	Flag     int16
	Vgroup   int16
	Protect  float32
	Uvname   [64]int8
}

// SDNA index: 112
type MultiresModifierData struct {
	Modifier  ModifierData
	Lvl       int8
	Sculptlvl int8
	Renderlvl int8
	Totlvl    int8
	Simple    int8
	Flags     int8
	Pad       [2]int8
}

// SDNA index: 113
type FluidsimModifierData struct {
	Modifier    ModifierData
	Fss         Pointer // *FluidsimSettings
	Point_cache Pointer // *PointCache
}

// SDNA index: 114
type ShrinkwrapModifierData struct {
	Modifier      ModifierData
	Target        Pointer // *Object
	AuxTarget     Pointer // *Object
	Vgroup_name   [64]int8
	KeepDist      float32
	ShrinkType    int16
	ShrinkOpts    int16
	ProjLimit     float32
	ProjAxis      int8
	SubsurfLevels int8
	Pad           [2]int8
}

// SDNA index: 115
type SimpleDeformModifierData struct {
	Modifier    ModifierData
	Origin      Pointer // *Object
	Vgroup_name [64]int8
	Factor      float32
	Limit       [2]float32
	Mode        int8
	Axis        int8
	OriginOpts  int8
	Pad         int8
}

// SDNA index: 116
type ShapeKeyModifierData struct {
	Modifier ModifierData
}

// SDNA index: 117
type SolidifyModifierData struct {
	Modifier      ModifierData
	Defgrp_name   [64]int8
	Offset        float32
	Offset_fac    float32
	Offset_fac_vg float32
	Offset_clamp  float32
	Pad           float32
	Crease_inner  float32
	Crease_outer  float32
	Crease_rim    float32
	Flag          int32
	Mat_ofs       int16
	Mat_ofs_rim   int16
}

// SDNA index: 118
type ScrewModifierData struct {
	Modifier     ModifierData
	Ob_axis      Pointer // *Object
	Steps        int32
	Render_steps int32
	Iter         int32
	Screw_ofs    float32
	Angle        float32
	Axis         int16
	Flag         int16
}

// SDNA index: 119
type OceanModifierData struct {
	Modifier       ModifierData
	Ocean          Pointer // *Ocean
	Oceancache     Pointer // *OceanCache
	Resolution     int32
	Spatial_size   int32
	Wind_velocity  float32
	Damp           float32
	Smallest_wave  float32
	Depth          float32
	Wave_alignment float32
	Wave_direction float32
	Wave_scale     float32
	Chop_amount    float32
	Foam_coverage  float32
	Time           float32
	Bakestart      int32
	Bakeend        int32
	Cachepath      [1024]int8
	Foamlayername  [64]int8
	Cached         int8
	Geometry_mode  int8
	Flag           int8
	Refresh        int8
	Repeat_x       int16
	Repeat_y       int16
	Seed           int32
	Size           float32
	Foam_fade      float32
	Pad            int32
}

// SDNA index: 120
type WarpModifierData struct {
	Modifier       ModifierData
	Texture        Pointer // *Tex
	Map_object     Pointer // *Object
	Uvlayer_name   [64]int8
	Uvlayer_tmp    int32
	Texmapping     int32
	Object_from    Pointer // *Object
	Object_to      Pointer // *Object
	Curfalloff     Pointer // *CurveMapping
	Defgrp_name    [64]int8
	Strength       float32
	Falloff_radius float32
	Flag           int8
	Falloff_type   int8
	Pad            [6]int8
}

// SDNA index: 121
type WeightVGEditModifierData struct {
	Modifier              ModifierData
	Defgrp_name           [64]int8
	Edit_flags            int16
	Falloff_type          int16
	Default_weight        float32
	Cmap_curve            Pointer // *CurveMapping
	Add_threshold         float32
	Rem_threshold         float32
	Mask_constant         float32
	Mask_defgrp_name      [64]int8
	Mask_tex_use_channel  int32
	Mask_texture          Pointer // *Tex
	Mask_tex_map_obj      Pointer // *Object
	Mask_tex_mapping      int32
	Mask_tex_uvlayer_name [64]int8
	Pad_i1                int32
}

// SDNA index: 122
type WeightVGMixModifierData struct {
	Modifier              ModifierData
	Defgrp_name_a         [64]int8
	Defgrp_name_b         [64]int8
	Default_weight_a      float32
	Default_weight_b      float32
	Mix_mode              int8
	Mix_set               int8
	Pad_c1                [6]int8
	Mask_constant         float32
	Mask_defgrp_name      [64]int8
	Mask_tex_use_channel  int32
	Mask_texture          Pointer // *Tex
	Mask_tex_map_obj      Pointer // *Object
	Mask_tex_mapping      int32
	Mask_tex_uvlayer_name [64]int8
	Pad_i1                int32
}

// SDNA index: 123
type WeightVGProximityModifierData struct {
	Modifier              ModifierData
	Defgrp_name           [64]int8
	Proximity_mode        int32
	Proximity_flags       int32
	Proximity_ob_target   Pointer // *Object
	Mask_constant         float32
	Mask_defgrp_name      [64]int8
	Mask_tex_use_channel  int32
	Mask_texture          Pointer // *Tex
	Mask_tex_map_obj      Pointer // *Object
	Mask_tex_mapping      int32
	Mask_tex_uvlayer_name [64]int8
	Min_dist              float32
	Max_dist              float32
	Falloff_type          int16
	Pad_s1                int16
}

// SDNA index: 124
type DynamicPaintModifierData struct {
	Modifier ModifierData
	Canvas   Pointer // *DynamicPaintCanvasSettings
	Brush    Pointer // *DynamicPaintBrushSettings
	Type     int32
	Pad      int32
}

// SDNA index: 125
type RemeshModifierData struct {
	Modifier    ModifierData
	Threshold   float32
	Scale       float32
	Hermite_num float32
	Depth       int8
	Flag        int8
	Mode        int8
	Pad         int8
}

// SDNA index: 126
type SkinModifierData struct {
	Modifier         ModifierData
	Branch_smoothing float32
	Flag             int8
	Symmetry_axes    int8
	Pad              [2]int8
}

// SDNA index: 127
type TriangulateModifierData struct {
	Modifier ModifierData
	Flag     int32
	Pad      int32
}

// SDNA index: 128
type LaplacianSmoothModifierData struct {
	Modifier      ModifierData
	Lambda        float32
	Lambda_border float32
	Pad1          float32
	Defgrp_name   [64]int8
	Flag          int16
	Repeat        int16
}

// SDNA index: 129
type UVWarpModifierData struct {
	Modifier     ModifierData
	Axis_u       int8
	Axis_v       int8
	Pad          [6]int8
	Center       [2]float32
	Object_src   Pointer // *Object
	Bone_src     [64]int8
	Object_dst   Pointer // *Object
	Bone_dst     [64]int8
	Vgroup_name  [64]int8
	Uvlayer_name [64]int8
}

// SDNA index: 130
type MeshCacheModifierData struct {
	Modifier     ModifierData
	Flag         int8
	Type         int8
	Time_mode    int8
	Play_mode    int8
	Forward_axis int8
	Up_axis      int8
	Flip_axis    int8
	Interp       int8
	Factor       float32
	Deform_mode  int8
	Pad          [7]int8
	Frame_start  float32
	Frame_scale  float32
	Eval_frame   float32
	Eval_time    float32
	Eval_factor  float32
	Filepath     [1024]int8
}

// SDNA index: 131
type EditLatt struct {
	Latt    Pointer // *Lattice
	Shapenr int32
	Pad     [4]int8
}

// SDNA index: 132
type Lattice struct {
	Id          ID
	Adt         Pointer // *AnimData
	Pntsu       int16
	Pntsv       int16
	Pntsw       int16
	Flag        int16
	Opntsu      int16
	Opntsv      int16
	Opntsw      int16
	Pad2        int16
	Typeu       int8
	Typev       int8
	Typew       int8
	Pad3        int8
	Pad         int32
	Fu          float32
	Fv          float32
	Fw          float32
	Du          float32
	Dv          float32
	Dw          float32
	Def         Pointer // *BPoint
	Ipo         Pointer // *Ipo
	Key         Pointer // *Key
	Dvert       Pointer // *MDeformVert
	Vgroup      [64]int8
	Latticedata Pointer // *float32
	Latmat      [4][4]float32
	Editlatt    Pointer // *EditLatt
}

// SDNA index: 133
type BDeformGroup struct {
	Next Pointer // *BDeformGroup
	Prev Pointer // *BDeformGroup
	Name [64]int8
	Flag int8
	Pad  [7]int8
}

// SDNA index: 134
type BoundBox struct {
	Vec  [8][3]float32
	Flag int32
	Pad  int32
}

// SDNA index: 135
type Object struct {
	Id                           ID
	Adt                          Pointer // *AnimData
	Sculpt                       Pointer // *SculptSession
	Type                         int16
	Partype                      int16
	Par1                         int32
	Par2                         int32
	Par3                         int32
	Parsubstr                    [64]int8
	Parent                       Pointer // *Object
	Track                        Pointer // *Object
	Proxy                        Pointer // *Object
	Proxy_group                  Pointer // *Object
	Proxy_from                   Pointer // *Object
	Ipo                          Pointer // *Ipo
	Bb                           Pointer // *BoundBox
	Action                       Pointer // *BAction
	Poselib                      Pointer // *BAction
	Pose                         Pointer // *BPose
	Data                         Pointer // *struct{}
	Gpd                          Pointer // *BGPdata
	Avs                          BAnimVizSettings
	Mpath                        Pointer // *BMotionPath
	ConstraintChannels           ListBase
	Effect                       ListBase
	Disp                         ListBase
	Defbase                      ListBase
	Modifiers                    ListBase
	Mode                         int32
	Restore_mode                 int32
	Mat                          Pointer // **Material
	Matbits                      Pointer // *int8
	Totcol                       int32
	Actcol                       int32
	Loc                          [3]float32
	Dloc                         [3]float32
	Orig                         [3]float32
	Size                         [3]float32
	Dsize                        [3]float32
	Dscale                       [3]float32
	Rot                          [3]float32
	Drot                         [3]float32
	Quat                         [4]float32
	Dquat                        [4]float32
	RotAxis                      [3]float32
	DrotAxis                     [3]float32
	RotAngle                     float32
	DrotAngle                    float32
	Obmat                        [4][4]float32
	Parentinv                    [4][4]float32
	Constinv                     [4][4]float32
	Imat                         [4][4]float32
	Imat_ren                     [4][4]float32
	Lay                          int32
	Sf                           float32
	Flag                         int16
	Colbits                      int16
	Transflag                    int16
	Protectflag                  int16
	Trackflag                    int16
	Upflag                       int16
	Nlaflag                      int16
	Ipoflag                      int16
	Scaflag                      int16
	Scavisflag                   int8
	Depsflag                     int8
	Dupon                        int32
	Dupoff                       int32
	Dupsta                       int32
	Dupend                       int32
	Mass                         float32
	Damping                      float32
	Inertia                      float32
	Formfactor                   float32
	Rdamping                     float32
	Sizefac                      float32
	Margin                       float32
	Max_vel                      float32
	Min_vel                      float32
	M_contactProcessingThreshold float32
	ObstacleRad                  float32
	Step_height                  float32
	Jump_speed                   float32
	Fall_speed                   float32
	Col_group                    int16
	Col_mask                     int16
	Rotmode                      int16
	Boundtype                    int8
	Collision_boundtype          int8
	Dtx                          int16
	Dt                           int8
	Empty_drawtype               int8
	Empty_drawsize               float32
	Dupfacesca                   float32
	Prop                         ListBase
	Sensors                      ListBase
	Controllers                  ListBase
	Actuators                    ListBase
	Bbsize                       [3]float32
	Index                        int16
	Actdef                       int16
	Col                          [4]float32
	Gameflag                     int32
	Gameflag2                    int32
	Bsoft                        Pointer // *BulletSoftBody
	Restrictflag                 int8
	Recalc                       int8
	Softflag                     int16
	AnisotropicFriction          [3]float32
	Constraints                  ListBase
	Nlastrips                    ListBase
	Hooks                        ListBase
	Particlesystem               ListBase
	Pd                           Pointer // *PartDeflect
	Soft                         Pointer // *SoftBody
	Dup_group                    Pointer // *Group
	Body_type                    int8
	Shapeflag                    int8
	Shapenr                      int16
	Smoothresh                   float32
	FluidsimSettings             Pointer // *FluidsimSettings
	DerivedDeform                Pointer // *DerivedMesh
	DerivedFinal                 Pointer // *DerivedMesh
	Pad                          Pointer // *int32
	LastDataMask                 uint64
	Customdata_mask              uint64
	State                        int32
	Init_state                   int32
	Gpulamp                      ListBase
	Pc_ids                       ListBase
	Duplilist                    Pointer // *ListBase
	Rigidbody_object             Pointer // *RigidBodyOb
	Rigidbody_constraint         Pointer // *RigidBodyCon
	Ima_ofs                      [2]float32
}

// SDNA index: 136
type ObHook struct {
	Next      Pointer // *ObHook
	Prev      Pointer // *ObHook
	Parent    Pointer // *Object
	Parentinv [4][4]float32
	Mat       [4][4]float32
	Cent      [3]float32
	Falloff   float32
	Name      [64]int8
	Indexar   Pointer // *int32
	Totindex  int32
	Curindex  int32
	Type      int16
	Active    int16
	Force     float32
}

// SDNA index: 137
type DupliObject struct {
	Next            Pointer // *DupliObject
	Prev            Pointer // *DupliObject
	Ob              Pointer // *Object
	Origlay         int32
	Pad             int32
	Mat             [4][4]float32
	Omat            [4][4]float32
	Orco            [3]float32
	Uv              [2]float32
	Type            int16
	No_draw         int8
	Animated        int8
	Persistent_id   [8]int32
	Particle_system Pointer // *ParticleSystem
}

// SDNA index: 138
type PartDeflect struct {
	Flag           int32
	Deflect        int16
	Forcefield     int16
	Falloff        int16
	Shape          int16
	Tex_mode       int16
	Kink           int16
	Kink_axis      int16
	Zdir           int16
	F_strength     float32
	F_damp         float32
	F_flow         float32
	F_size         float32
	F_power        float32
	Maxdist        float32
	Mindist        float32
	F_power_r      float32
	Maxrad         float32
	Minrad         float32
	Pdef_damp      float32
	Pdef_rdamp     float32
	Pdef_perm      float32
	Pdef_frict     float32
	Pdef_rfrict    float32
	Pdef_stickness float32
	Absorption     float32
	Pdef_sbdamp    float32
	Pdef_sbift     float32
	Pdef_sboft     float32
	Clump_fac      float32
	Clump_pow      float32
	Kink_freq      float32
	Kink_shape     float32
	Kink_amp       float32
	Free_end       float32
	Tex_nabla      float32
	Tex            Pointer // *Tex
	Rng            Pointer // *RNG
	F_noise        float32
	Seed           int32
	F_source       Pointer // *Object
}

// SDNA index: 139
type EffectorWeights struct {
	Group          Pointer // *Group
	Weight         [14]float32
	Global_gravity float32
	Flag           int16
	Rt             [3]int16
	Pad            int32
}

// SDNA index: 140
type PTCacheExtra struct {
	Next    Pointer // *PTCacheExtra
	Prev    Pointer // *PTCacheExtra
	Type    int32
	Totdata int32
	Data    Pointer // *struct{}
}

// SDNA index: 141
type PTCacheMem struct {
	Next       Pointer // *PTCacheMem
	Prev       Pointer // *PTCacheMem
	Frame      int32
	Totpoint   int32
	Data_types int32
	Flag       int32
	Data       [8]Pointer // [8]*struct{}
	Cur        [8]Pointer // [8]*struct{}
	Extradata  ListBase
}

// SDNA index: 142
type PointCache struct {
	Next          Pointer // *PointCache
	Prev          Pointer // *PointCache
	Flag          int32
	Step          int32
	Simframe      int32
	Startframe    int32
	Endframe      int32
	Editframe     int32
	Last_exact    int32
	Last_valid    int32
	Pad           int32
	Totpoint      int32
	Index         int32
	Compression   int16
	Rt            int16
	Name          [64]int8
	Prev_name     [64]int8
	Info          [64]int8
	Path          [1024]int8
	Cached_frames Pointer // *int8
	Mem_cache     ListBase
	Edit          Pointer // *PTCacheEdit
	Free_edit     func()
}

// SDNA index: 143
type SBVertex struct {
	Vec [4]float32
}

// SDNA index: 144
type BulletSoftBody struct {
	Flag                 int32
	LinStiff             float32
	AngStiff             float32
	Volume               float32
	Viterations          int32
	Piterations          int32
	Diterations          int32
	Citerations          int32
	KSRHR_CL             float32
	KSKHR_CL             float32
	KSSHR_CL             float32
	KSR_SPLT_CL          float32
	KSK_SPLT_CL          float32
	KSS_SPLT_CL          float32
	KVCF                 float32
	KDP                  float32
	KDG                  float32
	KLF                  float32
	KPR                  float32
	KVC                  float32
	KDF                  float32
	KMT                  float32
	KCHR                 float32
	KKHR                 float32
	KSHR                 float32
	KAHR                 float32
	Collisionflags       int32
	Numclusteriterations int32
	Welding              float32
	Margin               float32
}

// SDNA index: 145
type SoftBody struct {
	Totpoint         int32
	Totspring        int32
	Bpoint           Pointer // *BodyPoint
	Bspring          Pointer // *BodySpring
	Pad              int8
	Msg_lock         int8
	Msg_value        int16
	Nodemass         float32
	NamedVG_Mass     [64]int8
	Grav             float32
	Mediafrict       float32
	Rklimit          float32
	Physics_speed    float32
	Goalspring       float32
	Goalfrict        float32
	Mingoal          float32
	Maxgoal          float32
	Defgoal          float32
	Vertgroup        int16
	NamedVG_Softgoal [64]int8
	Fuzzyness        int16
	Inspring         float32
	Infrict          float32
	NamedVG_Spring_K [64]int8
	Sfra             int32
	Efra             int32
	Interval         int32
	Local            int16
	Solverflags      int16
	Keys             Pointer // **SBVertex
	Totpointkey      int32
	Totkey           int32
	Secondspring     float32
	Colball          float32
	Balldamp         float32
	Ballstiff        float32
	Sbc_mode         int16
	Aeroedge         int16
	Minloops         int16
	Maxloops         int16
	Choke            int16
	Solver_ID        int16
	Plastic          int16
	Springpreload    int16
	Scratch          Pointer // *SBScratch
	Shearstiff       float32
	Inpush           float32
	Pointcache       Pointer // *PointCache
	Ptcaches         ListBase
	Effector_weights Pointer // *EffectorWeights
	Lcom             [3]float32
	Lrot             [3][3]float32
	Lscale           [3][3]float32
	Last_frame       int32
}

// SDNA index: 146
type FluidVertexVelocity struct {
	Vel [3]float32
}

// SDNA index: 147
type FluidsimSettings struct {
	Fmd                   Pointer // *FluidsimModifierData
	Threads               int32
	Pad1                  int32
	Type                  int16
	Show_advancedoptions  int16
	Resolutionxyz         int16
	Previewresxyz         int16
	Realsize              float32
	GuiDisplayMode        int16
	RenderDisplayMode     int16
	ViscosityValue        float32
	ViscosityMode         int16
	ViscosityExponent     int16
	Grav                  [3]float32
	AnimStart             float32
	AnimEnd               float32
	BakeStart             int32
	BakeEnd               int32
	FrameOffset           int32
	Pad2                  int32
	Gstar                 float32
	MaxRefine             int32
	IniVelx               float32
	IniVely               float32
	IniVelz               float32
	OrgMesh               Pointer // *Mesh
	MeshBB                Pointer // *Mesh
	SurfdataPath          [1024]int8
	BbStart               [3]float32
	BbSize                [3]float32
	Ipo                   Pointer // *Ipo
	TypeFlags             int16
	DomainNovecgen        int8
	VolumeInitType        int8
	PartSlipValue         float32
	GenerateTracers       int32
	GenerateParticles     float32
	SurfaceSmoothing      float32
	SurfaceSubdivs        int32
	Flag                  int32
	ParticleInfSize       float32
	ParticleInfAlpha      float32
	FarFieldSize          float32
	MeshVelocities        Pointer // *FluidVertexVelocity
	Totvert               int32
	CpsTimeStart          float32
	CpsTimeEnd            float32
	CpsQuality            float32
	AttractforceStrength  float32
	AttractforceRadius    float32
	VelocityforceStrength float32
	VelocityforceRadius   float32
	Lastgoodframe         int32
	AnimRate              float32
}

// SDNA index: 148
type World struct {
	Id                   ID
	Adt                  Pointer // *AnimData
	Colormodel           int16
	Totex                int16
	Texact               int16
	Mistype              int16
	Horr                 float32
	Horg                 float32
	Horb                 float32
	Zenr                 float32
	Zeng                 float32
	Zenb                 float32
	Ambr                 float32
	Ambg                 float32
	Ambb                 float32
	Exposure             float32
	Exp                  float32
	Range                float32
	Linfac               float32
	Logfac               float32
	Gravity              float32
	ActivityBoxRadius    float32
	Skytype              int16
	Mode                 int16
	OcclusionRes         int16
	PhysicsEngine        int16
	Ticrate              int16
	Maxlogicstep         int16
	Physubstep           int16
	Maxphystep           int16
	Misi                 float32
	Miststa              float32
	Mistdist             float32
	Misthi               float32
	Starr                float32
	Starg                float32
	Starb                float32
	Stark                float32
	Starsize             float32
	Starmindist          float32
	Stardist             float32
	Starcolnoise         float32
	Dofsta               int16
	Dofend               int16
	Dofmin               int16
	Dofmax               int16
	Aodist               float32
	Aodistfac            float32
	Aoenergy             float32
	Aobias               float32
	Aomode               int16
	Aosamp               int16
	Aomix                int16
	Aocolor              int16
	Ao_adapt_thresh      float32
	Ao_adapt_speed_fac   float32
	Ao_approx_error      float32
	Ao_approx_correction float32
	Ao_indirect_energy   float32
	Ao_env_energy        float32
	Ao_pad2              float32
	Ao_indirect_bounces  int16
	Ao_pad               int16
	Ao_samp_method       int16
	Ao_gather_method     int16
	Ao_approx_passes     int16
	Flag                 int16
	Aosphere             Pointer     // *float32
	Aotables             Pointer     // *float32
	Ipo                  Pointer     // *Ipo
	Mtex                 [18]Pointer // [18]*MTex
	Pr_texture           int16
	Use_nodes            int16
	Pad                  [2]int16
	Preview              Pointer // *PreviewImage
	Nodetree             Pointer // *BNodeTree
}

// SDNA index: 149
type Base struct {
	Next   Pointer // *Base
	Prev   Pointer // *Base
	Lay    int32
	Selcol int32
	Flag   int32
	Sx     int16
	Sy     int16
	Object Pointer // *Object
}

// SDNA index: 150
type AviCodecData struct {
	LpFormat          Pointer // *struct{}
	LpParms           Pointer // *struct{}
	CbFormat          int32
	CbParms           int32
	FccType           int32
	FccHandler        int32
	DwKeyFrameEvery   int32
	DwQuality         int32
	DwBytesPerSecond  int32
	DwFlags           int32
	DwInterleaveEvery int32
	Pad               int32
	Avicodecname      [128]int8
}

// SDNA index: 151
type QuicktimeCodecData struct {
	CdParms     Pointer // *struct{}
	Pad         Pointer // *struct{}
	CdSize      int32
	Pad2        int32
	Qtcodecname [128]int8
}

// SDNA index: 152
type QuicktimeCodecSettings struct {
	CodecType            int32
	CodecSpatialQuality  int32
	Codec                int32
	CodecFlags           int32
	ColorDepth           int32
	CodecTemporalQuality int32
	MinSpatialQuality    int32
	MinTemporalQuality   int32
	KeyFrameRate         int32
	BitRate              int32
	AudiocodecType       int32
	AudioSampleRate      int32
	AudioBitDepth        int16
	AudioChannels        int16
	AudioCodecFlags      int32
	AudioBitRate         int32
	Pad1                 int32
}

// SDNA index: 153
type FFMpegCodecData struct {
	Type            int32
	Codec           int32
	Audio_codec     int32
	Video_bitrate   int32
	Audio_bitrate   int32
	Audio_mixrate   int32
	Audio_channels  int32
	Audio_pad       int32
	Audio_volume    float32
	Gop_size        int32
	Flags           int32
	Rc_min_rate     int32
	Rc_max_rate     int32
	Rc_buffer_size  int32
	Mux_packet_size int32
	Mux_rate        int32
	Properties      Pointer // *IDProperty
}

// SDNA index: 154
type AudioData struct {
	Mixrate        int32
	Main           float32
	Speed_of_sound float32
	Doppler_factor float32
	Distance_model int32
	Flag           int16
	Pad            int16
	Volume         float32
	Pad2           float32
}

// SDNA index: 155
type SceneRenderLayer struct {
	Next            Pointer // *SceneRenderLayer
	Prev            Pointer // *SceneRenderLayer
	Name            [64]int8
	Mat_override    Pointer // *Material
	Light_override  Pointer // *Group
	Lay             int32
	Lay_zmask       int32
	Lay_exclude     int32
	Layflag         int32
	Passflag        int32
	Pass_xor        int32
	Samples         int32
	Pad             int32
	FreestyleConfig FreestyleConfig
}

// SDNA index: 156
type ImageFormatData struct {
	Imtype           int8
	Depth            int8
	Planes           int8
	Flag             int8
	Quality          int8
	Compress         int8
	Exr_codec        int8
	Cineon_flag      int8
	Cineon_white     int16
	Cineon_black     int16
	Cineon_gamma     float32
	Jp2_flag         int8
	Jp2_codec        int8
	Pad              [6]int8
	View_settings    ColorManagedViewSettings
	Display_settings ColorManagedDisplaySettings
}

// SDNA index: 157
type RenderData struct {
	Im_format              ImageFormatData
	Avicodecdata           Pointer // *AviCodecData
	Qtcodecdata            Pointer // *QuicktimeCodecData
	Qtcodecsettings        QuicktimeCodecSettings
	Ffcodecdata            FFMpegCodecData
	Cfra                   int32
	Sfra                   int32
	Efra                   int32
	Subframe               float32
	Psfra                  int32
	Pefra                  int32
	Images                 int32
	Framapto               int32
	Flag                   int16
	Threads                int16
	Framelen               float32
	Blurfac                float32
	EdgeR                  float32
	EdgeG                  float32
	EdgeB                  float32
	Fullscreen             int16
	Xplay                  int16
	Yplay                  int16
	Freqplay               int16
	Depth                  int16
	Attrib                 int16
	Frame_step             int32
	Stereomode             int16
	Dimensionspreset       int16
	Filtertype             int16
	Size                   int16
	Maximsize              int16
	Pad6                   int16
	Xsch                   int32
	Ysch                   int32
	Xparts                 int16
	Yparts                 int16
	Tilex                  int32
	Tiley                  int32
	Planes                 int16
	Imtype                 int16
	Subimtype              int16
	Quality                int16
	Displaymode            int16
	Pad7                   int16
	Scemode                int32
	Mode                   int32
	Raytrace_options       int32
	Raytrace_structure     int16
	Pad1                   int16
	Ocres                  int16
	Pad4                   int16
	Alphamode              int16
	Osa                    int16
	Frs_sec                int16
	Edgeint                int16
	Safety                 Rctf
	Border                 Rctf
	Disprect               Rcti
	Layers                 ListBase
	Actlay                 int16
	Mblur_samples          int16
	Xasp                   float32
	Yasp                   float32
	Frs_sec_base           float32
	Gauss                  float32
	Color_mgt_flag         int32
	Postgamma              float32
	Posthue                float32
	Postsat                float32
	Dither_intensity       float32
	Bake_osa               int16
	Bake_filter            int16
	Bake_mode              int16
	Bake_flag              int16
	Bake_normal_space      int16
	Bake_quad_split        int16
	Bake_maxdist           float32
	Bake_biasdist          float32
	Bake_samples           int16
	Bake_pad               int16
	Pic                    [1024]int8
	Stamp                  int32
	Stamp_font_id          int16
	Pad3                   int16
	Stamp_udata            [768]int8
	Fg_stamp               [4]float32
	Bg_stamp               [4]float32
	Seq_prev_type          int8
	Seq_rend_type          int8
	Seq_flag               int8
	Pad5                   [5]int8
	Simplify_flag          int32
	Simplify_subsurf       int16
	Simplify_shadowsamples int16
	Simplify_particles     float32
	Simplify_aosss         float32
	Cineonwhite            int16
	Cineonblack            int16
	Cineongamma            float32
	Jp2_preset             int16
	Jp2_depth              int16
	Rpad3                  int32
	Domeres                int16
	Domemode               int16
	Domeangle              int16
	Dometilt               int16
	Domeresbuf             float32
	Pad2                   float32
	Dometext               Pointer // *Text
	Line_thickness_mode    int32
	Unit_line_thickness    float32
	Engine                 [32]int8
}

// SDNA index: 158
type RenderProfile struct {
	Next              Pointer // *RenderProfile
	Prev              Pointer // *RenderProfile
	Name              [32]int8
	Particle_perc     int16
	Subsurf_max       int16
	Shadbufsample_max int16
	Pad1              int16
	Ao_error          float32
	Pad2              float32
}

// SDNA index: 159
type GameDome struct {
	Res      int16
	Mode     int16
	Angle    int16
	Tilt     int16
	Resbuf   float32
	Pad2     float32
	Warptext Pointer // *Text
}

// SDNA index: 160
type GameFraming struct {
	Col  [3]float32
	Type int8
	Pad1 int8
	Pad2 int8
	Pad3 int8
}

// SDNA index: 161
type RecastData struct {
	Cellsize             float32
	Cellheight           float32
	Agentmaxslope        float32
	Agentmaxclimb        float32
	Agentheight          float32
	Agentradius          float32
	Edgemaxlen           float32
	Edgemaxerror         float32
	Regionminsize        float32
	Regionmergesize      float32
	Vertsperpoly         int32
	Detailsampledist     float32
	Detailsamplemaxerror float32
	Pad1                 int16
	Pad2                 int16
}

// SDNA index: 162
type GameData struct {
	Framing               GameFraming
	Playerflag            int16
	Xplay                 int16
	Yplay                 int16
	Freqplay              int16
	Depth                 int16
	Attrib                int16
	Rt1                   int16
	Rt2                   int16
	Aasamples             int16
	Pad4                  [3]int16
	Dome                  GameDome
	Stereoflag            int16
	Stereomode            int16
	Eyeseparation         float32
	RecastData            RecastData
	Gravity               float32
	ActivityBoxRadius     float32
	Flag                  int32
	Mode                  int16
	Matmode               int16
	OcclusionRes          int16
	PhysicsEngine         int16
	Exitkey               int16
	Pad                   int16
	Ticrate               int16
	Maxlogicstep          int16
	Physubstep            int16
	Maxphystep            int16
	ObstacleSimulation    int16
	Raster_storage        int16
	LevelHeight           float32
	Deactivationtime      float32
	Lineardeactthreshold  float32
	Angulardeactthreshold float32
	Pad2                  float32
}

// SDNA index: 163
type TimeMarker struct {
	Next   Pointer // *TimeMarker
	Prev   Pointer // *TimeMarker
	Frame  int32
	Name   [64]int8
	Flag   int32
	Camera Pointer // *Object
}

// SDNA index: 164
type Paint struct {
	Brush             Pointer // *Brush
	Paint_cursor      Pointer // *struct{}
	Paint_cursor_col  [4]int8
	Flags             int32
	Num_input_samples int32
	Pad               int32
}

// SDNA index: 165
type ImagePaintSettings struct {
	Paint            Paint
	Flag             int16
	Pad              int16
	Seam_bleed       int16
	Normal_angle     int16
	Screen_grab_size [2]int16
	Pad1             int32
	Paintcursor      Pointer // *struct{}
}

// SDNA index: 166
type ParticleBrushData struct {
	Size     int16
	Step     int16
	Invert   int16
	Count    int16
	Flag     int32
	Strength float32
}

// SDNA index: 167
type ParticleEditSettings struct {
	Flag        int16
	Totrekey    int16
	Totaddkey   int16
	Brushtype   int16
	Brush       [7]ParticleBrushData
	Paintcursor Pointer // *struct{}
	Emitterdist float32
	Rt          float32
	Selectmode  int32
	Edittype    int32
	Draw_step   int32
	Fade_frames int32
	Scene       Pointer // *Scene
	Object      Pointer // *Object
}

// SDNA index: 168
type Sculpt struct {
	Paint                Paint
	Flags                int32
	Radial_symm          [3]int32
	Detail_size          int32
	Symmetrize_direction int32
}

// SDNA index: 169
type UvSculpt struct {
	Paint Paint
}

// SDNA index: 170
type VPaint struct {
	Paint       Paint
	Flag        int16
	Pad         int16
	Tot         int32
	Vpaint_prev Pointer // *int32
	Wpaint_prev Pointer // *MDeformVert
	Paintcursor Pointer // *struct{}
}

// SDNA index: 171
type TransformOrientation struct {
	Next Pointer // *TransformOrientation
	Prev Pointer // *TransformOrientation
	Name [64]int8
	Mat  [3][3]float32
	Pad  int32
}

// SDNA index: 172
type UnifiedPaintSettings struct {
	Size                   int32
	Unprojected_radius     float32
	Alpha                  float32
	Weight                 float32
	Flag                   int32
	Last_rake              [2]float32
	Pad                    int32
	Brush_rotation         float32
	Draw_anchored          int32
	Anchored_size          int32
	Anchored_initial_mouse [2]float32
	Draw_pressure          int32
	Pressure_value         float32
	Tex_mouse              [2]float32
	Mask_tex_mouse         [2]float32
	Pixel_radius           float32
}

// SDNA index: 173
type MeshStatVis struct {
	Type              int8
	_pad1             [2]int8
	Overhang_axis     int8
	Overhang_min      float32
	Overhang_max      float32
	Thickness_min     float32
	Thickness_max     float32
	Thickness_samples int8
	_pad2             [3]int8
	Distort_min       float32
	Distort_max       float32
	Sharp_min         float32
	Sharp_max         float32
}

// SDNA index: 174
type ToolSettings struct {
	Vpaint                                  Pointer // *VPaint
	Wpaint                                  Pointer // *VPaint
	Sculpt                                  Pointer // *Sculpt
	Uvsculpt                                Pointer // *UvSculpt
	Vgroup_weight                           float32
	Cornertype                              int16
	Pad1                                    int16
	Jointrilimit                            float32
	Degr                                    float32
	Step                                    int16
	Turn                                    int16
	Extr_offs                               float32
	Doublimit                               float32
	Normalsize                              float32
	Automerge                               int16
	Selectmode                              int16
	Segments                                int16
	Rings                                   int16
	Vertices                                int16
	Unwrapper                               int16
	Uvcalc_radius                           float32
	Uvcalc_cubesize                         float32
	Uvcalc_margin                           float32
	Uvcalc_mapdir                           int16
	Uvcalc_mapalign                         int16
	Uvcalc_flag                             int16
	Uv_flag                                 int16
	Uv_selectmode                           int16
	Pad2                                    int16
	Gpencil_flags                           int16
	Autoik_chainlen                         int16
	Imapaint                                ImagePaintSettings
	Particle                                ParticleEditSettings
	Proportional_size                       float32
	Select_thresh                           float32
	Clean_thresh                            float32
	Autokey_mode                            int16
	Autokey_flag                            int16
	Multires_subdiv_type                    int8
	Pad3                                    [5]int8
	Skgen_resolution                        int16
	Skgen_threshold_internal                float32
	Skgen_threshold_external                float32
	Skgen_length_ratio                      float32
	Skgen_length_limit                      float32
	Skgen_angle_limit                       float32
	Skgen_correlation_limit                 float32
	Skgen_symmetry_limit                    float32
	Skgen_retarget_angle_weight             float32
	Skgen_retarget_length_weight            float32
	Skgen_retarget_distance_weight          float32
	Skgen_options                           int16
	Skgen_postpro                           int8
	Skgen_postpro_passes                    int8
	Skgen_subdivisions                      [3]int8
	Skgen_multi_level                       int8
	Skgen_template                          Pointer // *Object
	Bone_sketching                          int8
	Bone_sketching_convert                  int8
	Skgen_subdivision_number                int8
	Skgen_retarget_options                  int8
	Skgen_retarget_roll                     int8
	Skgen_side_string                       [8]int8
	Skgen_num_string                        [8]int8
	Edge_mode                               int8
	Edge_mode_live_unwrap                   int8
	Snap_mode                               int8
	Snap_node_mode                          int8
	Snap_uv_mode                            int8
	Snap_flag                               int16
	Snap_target                             int16
	Proportional                            int16
	Prop_mode                               int16
	Proportional_objects                    int8
	Proportional_mask                       int8
	Pad4                                    [1]int8
	Auto_normalize                          int8
	Multipaint                              int8
	Weightuser                              int8
	Use_uv_sculpt                           int32
	Uv_sculpt_settings                      int32
	Uv_sculpt_tool                          int32
	Uv_relax_method                         int32
	Sculpt_paint_settings                   int16
	Pad5                                    int16
	Sculpt_paint_unified_size               int32
	Sculpt_paint_unified_unprojected_radius float32
	Sculpt_paint_unified_alpha              float32
	Unified_paint_settings                  UnifiedPaintSettings
	Statvis                                 MeshStatVis
}

// SDNA index: 175
type BStats struct {
	Totobj      int32
	Totlamp     int32
	Totobjsel   int32
	Totcurve    int32
	Totmesh     int32
	Totarmature int32
	Totvert     int32
	Totface     int32
}

// SDNA index: 176
type UnitSettings struct {
	Scale_length    float32
	System          int8
	System_rotation int8
	Flag            int16
}

// SDNA index: 177
type PhysicsSettings struct {
	Gravity          [3]float32
	Flag             int32
	Quick_cache_step int32
	Rt               int32
}

// SDNA index: 178
type Scene struct {
	Id                            ID
	Adt                           Pointer // *AnimData
	Camera                        Pointer // *Object
	World                         Pointer // *World
	Set                           Pointer // *Scene
	Base                          ListBase
	Basact                        Pointer // *Base
	Obedit                        Pointer // *Object
	Cursor                        [3]float32
	Twcent                        [3]float32
	Twmin                         [3]float32
	Twmax                         [3]float32
	Lay                           int32
	Layact                        int32
	Lay_updated                   int32
	Flag                          int16
	Use_nodes                     int16
	Nodetree                      Pointer // *BNodeTree
	Ed                            Pointer // *Editing
	Toolsettings                  Pointer // *ToolSettings
	Stats                         Pointer // *SceneStats
	R                             RenderData
	Audio                         AudioData
	Markers                       ListBase
	Transform_spaces              ListBase
	Sound_scene                   Pointer // *struct{}
	Sound_scene_handle            Pointer // *struct{}
	Sound_scrub_handle            Pointer // *struct{}
	Speaker_handles               Pointer // *struct{}
	Fps_info                      Pointer // *struct{}
	TheDag                        Pointer // *DagForest
	Dagflags                      int16
	Recalc                        int16
	Active_keyingset              int32
	Keyingsets                    ListBase
	Framing                       GameFraming
	Gm                            GameData
	Unit                          UnitSettings
	Gpd                           Pointer // *BGPdata
	Physics_settings              PhysicsSettings
	Clip                          Pointer // *MovieClip
	Customdata_mask               uint64
	Customdata_mask_modal         uint64
	View_settings                 ColorManagedViewSettings
	Display_settings              ColorManagedDisplaySettings
	Sequencer_colorspace_settings ColorManagedColorspaceSettings
	Rigidbody_world               Pointer // *RigidBodyWorld
}

// SDNA index: 179
type BGpic struct {
	Next   Pointer // *BGpic
	Prev   Pointer // *BGpic
	Ima    Pointer // *Image
	Iuser  ImageUser
	Clip   Pointer // *MovieClip
	Cuser  MovieClipUser
	Xof    float32
	Yof    float32
	Size   float32
	Blend  float32
	View   int16
	Flag   int16
	Source int16
	Pad    int16
}

// SDNA index: 180
type RegionView3D struct {
	Winmat        [4][4]float32
	Viewmat       [4][4]float32
	Viewinv       [4][4]float32
	Persmat       [4][4]float32
	Persinv       [4][4]float32
	Viewmatob     [4][4]float32
	Persmatob     [4][4]float32
	Clip          [6][4]float32
	Clip_local    [6][4]float32
	Clipbb        Pointer // *BoundBox
	Gpd           Pointer // *BGPdata
	Localvd       Pointer // *RegionView3D
	Ri            Pointer // *RenderInfo
	Render_engine Pointer // *RenderEngine
	Depths        Pointer // *ViewDepths
	Gpuoffscreen  Pointer // *struct{}
	Sms           Pointer // *SmoothView3DStore
	Smooth_timer  Pointer // *WmTimer
	Twmat         [4][4]float32
	Viewquat      [4]float32
	Dist          float32
	Camdx         float32
	Camdy         float32
	Pixsize       float32
	Ofs           [3]float32
	Camzoom       float32
	Is_persp      int8
	Persp         int8
	View          int8
	Viewlock      int8
	Pad           [4]int8
	Twdrawflag    int16
	Rflag         int16
	Lviewquat     [4]float32
	Lpersp        int16
	Lview         int16
	Gridview      float32
	Twangle       [3]float32
	Rot_angle     float32
	Rot_axis      [3]float32
}

// SDNA index: 181
type View3D struct {
	Next                 Pointer // *SpaceLink
	Prev                 Pointer // *SpaceLink
	Regionbase           ListBase
	Spacetype            int32
	Blockscale           float32
	Blockhandler         [8]int16
	Viewquat             [4]float32
	Dist                 float32
	Bundle_size          float32
	Bundle_drawtype      int16
	Pad                  int16
	Matcap_icon          int32
	Lay_used             int32
	Persp                int16
	View                 int16
	Camera               Pointer // *Object
	Ob_centre            Pointer // *Object
	Render_border        Rctf
	Bgpicbase            ListBase
	Bgpic                Pointer // *BGpic
	Localvd              Pointer // *View3D
	Ob_centre_bone       [64]int8
	Lay                  int32
	Layact               int32
	Drawtype             int16
	Ob_centre_cursor     int16
	Scenelock            int16
	Around               int16
	Flag                 int16
	Flag2                int16
	Lens                 float32
	Grid                 float32
	Near                 float32
	Far                  float32
	Ofs                  [3]float32
	Cursor               [3]float32
	Modeselect           int16
	Gridlines            int16
	Gridsubdiv           int16
	Gridflag             int8
	Twtype               int8
	Twmode               int8
	Twflag               int8
	Pad2                 [2]int8
	Afterdraw_transp     ListBase
	Afterdraw_xray       ListBase
	Afterdraw_xraytransp ListBase
	Zbuf                 int16
	Transp               int16
	Xray                 int16
	Pad3                 [2]int8
	Properties_storage   Pointer // *struct{}
	Defmaterial          Pointer // *Material
	Gpd                  Pointer // *BGPdata
}

// SDNA index: 182
type View2D struct {
	Tot          Rctf
	Cur          Rctf
	Vert         Rcti
	Hor          Rcti
	Mask         Rcti
	Min          [2]float32
	Max          [2]float32
	Minzoom      float32
	Maxzoom      float32
	Scroll       int16
	Scroll_ui    int16
	Keeptot      int16
	Keepzoom     int16
	Keepofs      int16
	Flag         int16
	Align        int16
	Winx         int16
	Winy         int16
	Oldwinx      int16
	Oldwiny      int16
	Around       int16
	Tab_offset   Pointer // *float32
	Tab_num      int32
	Tab_cur      int32
	Sms          Pointer // *SmoothView2DStore
	Smooth_timer Pointer // *WmTimer
}

// SDNA index: 183
type SpaceLink struct {
	Next         Pointer // *SpaceLink
	Prev         Pointer // *SpaceLink
	Regionbase   ListBase
	Spacetype    int32
	Blockscale   float32
	Blockhandler [8]int16
}

// SDNA index: 184
type SpaceInfo struct {
	Next         Pointer // *SpaceLink
	Prev         Pointer // *SpaceLink
	Regionbase   ListBase
	Spacetype    int32
	Blockscale   float32
	Blockhandler [8]int16
	Rpt_mask     int8
	Pad          [7]int8
}

// SDNA index: 185
type SpaceButs struct {
	Next            Pointer // *SpaceLink
	Prev            Pointer // *SpaceLink
	Regionbase      ListBase
	Spacetype       int32
	Blockscale      float32
	Blockhandler    [8]int16
	V2d             View2D
	Mainb           int16
	Mainbo          int16
	Mainbuser       int16
	Re_align        int16
	Align           int16
	Preview         int16
	Texture_context int16
	Flag            int8
	Pad             int8
	Path            Pointer // *struct{}
	Pathflag        int32
	Dataicon        int32
	Pinid           Pointer // *ID
	Texuser         Pointer // *struct{}
}

// SDNA index: 186
type SpaceOops struct {
	Next          Pointer // *SpaceLink
	Prev          Pointer // *SpaceLink
	Regionbase    ListBase
	Spacetype     int32
	Blockscale    float32
	Blockhandler  [8]int16
	V2d           View2D
	Tree          ListBase
	Treestore     Pointer // *TreeStore
	Search_string [32]int8
	Search_tse    TreeStoreElem
	Flag          int16
	Outlinevis    int16
	Storeflag     int16
	Search_flags  int16
}

// SDNA index: 187
type SpaceIpo struct {
	Next         Pointer // *SpaceLink
	Prev         Pointer // *SpaceLink
	Regionbase   ListBase
	Spacetype    int32
	Blockscale   float32
	Blockhandler [8]int16
	V2d          View2D
	Ads          Pointer // *BDopeSheet
	GhostCurves  ListBase
	Mode         int16
	Autosnap     int16
	Flag         int32
	CursorVal    float32
	Around       int32
}

// SDNA index: 188
type SpaceNla struct {
	Next         Pointer // *SpaceLink
	Prev         Pointer // *SpaceLink
	Regionbase   ListBase
	Spacetype    int32
	Blockscale   float32
	Blockhandler [8]int16
	Autosnap     int16
	Flag         int16
	Pad          int32
	Ads          Pointer // *BDopeSheet
	V2d          View2D
}

// SDNA index: 189
type SpaceTimeCache struct {
	Next  Pointer // *SpaceTimeCache
	Prev  Pointer // *SpaceTimeCache
	Array Pointer // *float32
}

// SDNA index: 190
type SpaceTime struct {
	Next          Pointer // *SpaceLink
	Prev          Pointer // *SpaceLink
	Regionbase    ListBase
	Spacetype     int32
	Blockscale    float32
	V2d           View2D
	Caches        ListBase
	Cache_display int32
	Flag          int32
}

// SDNA index: 191
type SpaceSeq struct {
	Next         Pointer // *SpaceLink
	Prev         Pointer // *SpaceLink
	Regionbase   ListBase
	Spacetype    int32
	Blockscale   float32
	Blockhandler [8]int16
	V2d          View2D
	Xof          float32
	Yof          float32
	Mainb        int16
	Render_size  int16
	Chanshown    int16
	Zebra        int16
	Flag         int32
	Zoom         float32
	View         int32
	Overlay_type int32
	Gpd          Pointer // *BGPdata
	Scopes       SequencerScopes
}

// SDNA index: 192
type MaskSpaceInfo struct {
	Mask      Pointer // *Mask
	Draw_flag int8
	Draw_type int8
	Pad3      [6]int8
}

// SDNA index: 193
type FileSelectParams struct {
	Title       [96]int8
	Dir         [1056]int8
	File        [256]int8
	Renamefile  [256]int8
	Renameedit  [256]int8
	Filter_glob [64]int8
	Active_file int32
	Sel_first   int32
	Sel_last    int32
	Type        int16
	Flag        int16
	Sort        int16
	Display     int16
	Filter      int16
	F_fp        int16
	Fp_str      [8]int8
}

// SDNA index: 194
type SpaceFile struct {
	Next               Pointer // *SpaceLink
	Prev               Pointer // *SpaceLink
	Regionbase         ListBase
	Spacetype          int32
	Scroll_offset      int32
	Params             Pointer // *FileSelectParams
	Files              Pointer // *FileList
	Folders_prev       Pointer // *ListBase
	Folders_next       Pointer // *ListBase
	Op                 Pointer // *WmOperator
	Smoothscroll_timer Pointer // *WmTimer
	Layout             Pointer // *FileLayout
	Recentnr           int16
	Bookmarknr         int16
	Systemnr           int16
	Pad2               int16
}

// SDNA index: 195
type SpaceImage struct {
	Next             Pointer // *SpaceLink
	Prev             Pointer // *SpaceLink
	Regionbase       ListBase
	Spacetype        int32
	Flag             int32
	Image            Pointer // *Image
	Iuser            ImageUser
	Cumap            Pointer // *CurveMapping
	Scopes           Scopes
	Sample_line_hist Histogram
	Gpd              Pointer // *BGPdata
	Cursor           [2]float32
	Xof              float32
	Yof              float32
	Zoom             float32
	Centx            float32
	Centy            float32
	Mode             int8
	Pin              int8
	Pad              int16
	Curtile          int16
	Lock             int16
	Dt_uv            int8
	Sticky           int8
	Dt_uvstretch     int8
	Around           int8
	Mask_info        MaskSpaceInfo
}

// SDNA index: 196
type SpaceText struct {
	Next          Pointer // *SpaceLink
	Prev          Pointer // *SpaceLink
	Regionbase    ListBase
	Spacetype     int32
	Blockscale    float32
	Blockhandler  [8]int16
	Text          Pointer // *Text
	Top           int32
	Viewlines     int32
	Flags         int16
	Menunr        int16
	Lheight       int16
	Cwidth        int8
	Linenrs_tot   int8
	Left          int32
	Showlinenrs   int32
	Tabnumber     int32
	Showsyntax    int16
	Line_hlight   int16
	Overwrite     int16
	Live_edit     int16
	Pix_per_line  float32
	Txtscroll     Rcti
	Txtbar        Rcti
	Wordwrap      int32
	Doplugins     int32
	Findstr       [256]int8
	Replacestr    [256]int8
	Margin_column int16
	Lheight_dpi   int16
	Pad           [4]int8
	Drawcache     Pointer // *struct{}
}

// SDNA index: 197
type Script struct {
	Id                 ID
	Py_draw            Pointer // *struct{}
	Py_event           Pointer // *struct{}
	Py_button          Pointer // *struct{}
	Py_browsercallback Pointer // *struct{}
	Py_globaldict      Pointer // *struct{}
	Flags              int32
	Lastspace          int32
	Scriptname         [1024]int8
	Scriptarg          [256]int8
}

// SDNA index: 198
type SpaceScript struct {
	Next       Pointer // *SpaceLink
	Prev       Pointer // *SpaceLink
	Regionbase ListBase
	Spacetype  int32
	Blockscale float32
	Script     Pointer // *Script
	Flags      int16
	Menunr     int16
	Pad1       int32
	But_refs   Pointer // *struct{}
}

// SDNA index: 199
type BNodeTreePath struct {
	Next        Pointer // *BNodeTreePath
	Prev        Pointer // *BNodeTreePath
	Nodetree    Pointer // *BNodeTree
	Parent_key  BNodeInstanceKey
	Pad         int32
	View_center [2]float32
	Node_name   [64]int8
}

// SDNA index: 200
type SpaceNode struct {
	Next         Pointer // *SpaceLink
	Prev         Pointer // *SpaceLink
	Regionbase   ListBase
	Spacetype    int32
	Blockscale   float32
	Blockhandler [8]int16
	V2d          View2D
	Id           Pointer // *ID
	From         Pointer // *ID
	Flag         int16
	Pad1         int16
	Aspect       float32
	Pad2         float32
	Xof          float32
	Yof          float32
	Zoom         float32
	Cursor       [2]float32
	Treepath     ListBase
	Nodetree     Pointer // *BNodeTree
	Edittree     Pointer // *BNodeTree
	Tree_idname  [64]int8
	Treetype     int32
	Pad3         int32
	Texfrom      int16
	Shaderfrom   int16
	Recalc       int16
	Pad4         int16
	Linkdrag     ListBase
	Gpd          Pointer // *BGPdata
}

// SDNA index: 201
type SpaceLogic struct {
	Next         Pointer // *SpaceLink
	Prev         Pointer // *SpaceLink
	Regionbase   ListBase
	Spacetype    int32
	Blockscale   float32
	Blockhandler [8]int16
	Flag         int16
	Scaflag      int16
	Pad          int32
	Gpd          Pointer // *BGPdata
}

// SDNA index: 202
type ConsoleLine struct {
	Next      Pointer // *ConsoleLine
	Prev      Pointer // *ConsoleLine
	Len_alloc int32
	Len       int32
	Line      Pointer // *int8
	Cursor    int32
	Type      int32
}

// SDNA index: 203
type SpaceConsole struct {
	Next         Pointer // *SpaceLink
	Prev         Pointer // *SpaceLink
	Regionbase   ListBase
	Spacetype    int32
	Blockscale   float32
	Blockhandler [8]int16
	Lheight      int32
	Pad          int32
	Scrollback   ListBase
	History      ListBase
	Prompt       [256]int8
	Language     [32]int8
	Sel_start    int32
	Sel_end      int32
}

// SDNA index: 204
type SpaceUserPref struct {
	Next        Pointer // *SpaceLink
	Prev        Pointer // *SpaceLink
	Regionbase  ListBase
	Spacetype   int32
	Pad         [3]int8
	Filter_type int8
	Filter      [64]int8
}

// SDNA index: 205
type SpaceClip struct {
	Next          Pointer // *SpaceLink
	Prev          Pointer // *SpaceLink
	Regionbase    ListBase
	Spacetype     int32
	Xof           float32
	Yof           float32
	Xlockof       float32
	Ylockof       float32
	Zoom          float32
	User          MovieClipUser
	Clip          Pointer // *MovieClip
	Scopes        MovieClipScopes
	Flag          int32
	Mode          int16
	View          int16
	Path_length   int32
	Loc           [2]float32
	Scale         float32
	Angle         float32
	Pad           int32
	Stabmat       [4][4]float32
	Unistabmat    [4][4]float32
	Postproc_flag int32
	Gpencil_src   int16
	Pad2          int16
	Around        int32
	Pad4          int32
	Mask_info     MaskSpaceInfo
}

// SDNA index: 206
type UiFont struct {
	Next      Pointer // *UiFont
	Prev      Pointer // *UiFont
	Filename  [1024]int8
	Blf_id    int16
	Uifont_id int16
	R_to_l    int16
	Hinting   int16
}

// SDNA index: 207
type UiFontStyle struct {
	Uifont_id   int16
	Points      int16
	Kerning     int16
	Pad         [6]int8
	Italic      int16
	Bold        int16
	Shadow      int16
	Shadx       int16
	Shady       int16
	Align       int16
	Shadowalpha float32
	Shadowcolor float32
}

// SDNA index: 208
type UiStyle struct {
	Next           Pointer // *UiStyle
	Prev           Pointer // *UiStyle
	Name           [64]int8
	Paneltitle     UiFontStyle
	Grouplabel     UiFontStyle
	Widgetlabel    UiFontStyle
	Widget         UiFontStyle
	Panelzoom      float32
	Minlabelchars  int16
	Minwidgetchars int16
	Columnspace    int16
	Templatespace  int16
	Boxspace       int16
	Buttonspacex   int16
	Buttonspacey   int16
	Panelspace     int16
	Panelouter     int16
	Pad            int16
}

// SDNA index: 209
type UiWidgetColors struct {
	Outline     [4]int8
	Inner       [4]int8
	Inner_sel   [4]int8
	Item        [4]int8
	Text        [4]int8
	Text_sel    [4]int8
	Shaded      int16
	Shadetop    int16
	Shadedown   int16
	Alpha_check int16
}

// SDNA index: 210
type UiWidgetStateColors struct {
	Inner_anim       [4]int8
	Inner_anim_sel   [4]int8
	Inner_key        [4]int8
	Inner_key_sel    [4]int8
	Inner_driven     [4]int8
	Inner_driven_sel [4]int8
	Blend            float32
	Pad              float32
}

// SDNA index: 211
type UiPanelColors struct {
	Header      [4]int8
	Back        [4]int8
	Show_header int16
	Show_back   int16
	Pad         int32
}

// SDNA index: 212
type UiGradientColors struct {
	Gradient      [4]int8
	High_gradient [4]int8
	Show_grad     int32
	Pad2          int32
}

// SDNA index: 213
type ThemeUI struct {
	Wcol_regular      UiWidgetColors
	Wcol_tool         UiWidgetColors
	Wcol_text         UiWidgetColors
	Wcol_radio        UiWidgetColors
	Wcol_option       UiWidgetColors
	Wcol_toggle       UiWidgetColors
	Wcol_num          UiWidgetColors
	Wcol_numslider    UiWidgetColors
	Wcol_menu         UiWidgetColors
	Wcol_pulldown     UiWidgetColors
	Wcol_menu_back    UiWidgetColors
	Wcol_menu_item    UiWidgetColors
	Wcol_tooltip      UiWidgetColors
	Wcol_box          UiWidgetColors
	Wcol_scroll       UiWidgetColors
	Wcol_progress     UiWidgetColors
	Wcol_list_item    UiWidgetColors
	Wcol_state        UiWidgetStateColors
	Panel             UiPanelColors
	Menu_shadow_fac   float32
	Menu_shadow_width int16
	Pad               int16
	Iconfile          [256]int8
	Icon_alpha        float32
	Xaxis             [4]int8
	Yaxis             [4]int8
	Zaxis             [4]int8
}

// SDNA index: 214
type ThemeSpace struct {
	Back                        [4]int8
	Title                       [4]int8
	Text                        [4]int8
	Text_hi                     [4]int8
	Header                      [4]int8
	Header_title                [4]int8
	Header_text                 [4]int8
	Header_text_hi              [4]int8
	Button                      [4]int8
	Button_title                [4]int8
	Button_text                 [4]int8
	Button_text_hi              [4]int8
	List                        [4]int8
	List_title                  [4]int8
	List_text                   [4]int8
	List_text_hi                [4]int8
	Panelcolors                 UiPanelColors
	Gradients                   UiGradientColors
	Shade1                      [4]int8
	Shade2                      [4]int8
	Hilite                      [4]int8
	Grid                        [4]int8
	Wire                        [4]int8
	Select                      [4]int8
	Lamp                        [4]int8
	Speaker                     [4]int8
	Empty                       [4]int8
	Camera                      [4]int8
	Pad                         [8]int8
	Active                      [4]int8
	Group                       [4]int8
	Group_active                [4]int8
	Transform                   [4]int8
	Vertex                      [4]int8
	Vertex_select               [4]int8
	Vertex_unreferenced         [4]int8
	Edge                        [4]int8
	Edge_select                 [4]int8
	Edge_seam                   [4]int8
	Edge_sharp                  [4]int8
	Edge_facesel                [4]int8
	Edge_crease                 [4]int8
	Face                        [4]int8
	Face_select                 [4]int8
	Face_dot                    [4]int8
	Extra_edge_len              [4]int8
	Extra_edge_angle            [4]int8
	Extra_face_angle            [4]int8
	Extra_face_area             [4]int8
	Normal                      [4]int8
	Vertex_normal               [4]int8
	Bone_solid                  [4]int8
	Bone_pose                   [4]int8
	Bone_pose_active            [4]int8
	Strip                       [4]int8
	Strip_select                [4]int8
	Cframe                      [4]int8
	Freestyle_edge_mark         [4]int8
	Freestyle_face_mark         [4]int8
	Nurb_uline                  [4]int8
	Nurb_vline                  [4]int8
	Act_spline                  [4]int8
	Nurb_sel_uline              [4]int8
	Nurb_sel_vline              [4]int8
	Lastsel_point               [4]int8
	Handle_free                 [4]int8
	Handle_auto                 [4]int8
	Handle_vect                 [4]int8
	Handle_align                [4]int8
	Handle_auto_clamped         [4]int8
	Handle_sel_free             [4]int8
	Handle_sel_auto             [4]int8
	Handle_sel_vect             [4]int8
	Handle_sel_align            [4]int8
	Handle_sel_auto_clamped     [4]int8
	Ds_channel                  [4]int8
	Ds_subchannel               [4]int8
	Console_output              [4]int8
	Console_input               [4]int8
	Console_info                [4]int8
	Console_error               [4]int8
	Console_cursor              [4]int8
	Console_select              [4]int8
	Pad1                        [4]int8
	Vertex_size                 int8
	Outline_width               int8
	Facedot_size                int8
	Noodle_curving              int8
	Syntaxl                     [4]int8
	Syntaxs                     [4]int8
	Syntaxb                     [4]int8
	Syntaxn                     [4]int8
	Syntaxv                     [4]int8
	Syntaxc                     [4]int8
	Syntaxd                     [4]int8
	Syntaxr                     [4]int8
	Movie                       [4]int8
	Movieclip                   [4]int8
	Mask                        [4]int8
	Image                       [4]int8
	Scene                       [4]int8
	Audio                       [4]int8
	Effect                      [4]int8
	Transition                  [4]int8
	Meta                        [4]int8
	Editmesh_active             [4]int8
	Handle_vertex               [4]int8
	Handle_vertex_select        [4]int8
	Pad2                        [4]int8
	Handle_vertex_size          int8
	Marker_outline              [4]int8
	Marker                      [4]int8
	Act_marker                  [4]int8
	Sel_marker                  [4]int8
	Dis_marker                  [4]int8
	Lock_marker                 [4]int8
	Bundle_solid                [4]int8
	Path_before                 [4]int8
	Path_after                  [4]int8
	Camera_path                 [4]int8
	Hpad                        [3]int8
	Preview_back                [4]int8
	Preview_stitch_face         [4]int8
	Preview_stitch_edge         [4]int8
	Preview_stitch_vert         [4]int8
	Preview_stitch_stitchable   [4]int8
	Preview_stitch_unstitchable [4]int8
	Preview_stitch_active       [4]int8
	Match                       [4]int8
	Selected_highlight          [4]int8
	Skin_root                   [4]int8
	Anim_active                 [4]int8
	Anim_non_active             [4]int8
	Nla_tweaking                [4]int8
	Nla_tweakdupli              [4]int8
	Nla_transition              [4]int8
	Nla_transition_sel          [4]int8
	Nla_meta                    [4]int8
	Nla_meta_sel                [4]int8
	Nla_sound                   [4]int8
	Nla_sound_sel               [4]int8
}

// SDNA index: 215
type ThemeWireColor struct {
	Solid  [4]int8
	Select [4]int8
	Active [4]int8
	Flag   int16
	Pad    int16
}

// SDNA index: 216
type BTheme struct {
	Next              Pointer // *BTheme
	Prev              Pointer // *BTheme
	Name              [32]int8
	Tui               ThemeUI
	Tbuts             ThemeSpace
	Tv3d              ThemeSpace
	Tfile             ThemeSpace
	Tipo              ThemeSpace
	Tinfo             ThemeSpace
	Tact              ThemeSpace
	Tnla              ThemeSpace
	Tseq              ThemeSpace
	Tima              ThemeSpace
	Text              ThemeSpace
	Toops             ThemeSpace
	Ttime             ThemeSpace
	Tnode             ThemeSpace
	Tlogic            ThemeSpace
	Tuserpref         ThemeSpace
	Tconsole          ThemeSpace
	Tclip             ThemeSpace
	Tarm              [20]ThemeWireColor
	Active_theme_area int32
	Pad               int32
}

// SDNA index: 217
type BAddon struct {
	Next   Pointer // *BAddon
	Prev   Pointer // *BAddon
	Module [64]int8
	Prop   Pointer // *IDProperty
}

// SDNA index: 218
type SolidLight struct {
	Flag int32
	Pad  int32
	Col  [4]float32
	Spec [4]float32
	Vec  [4]float32
}

// SDNA index: 219
type UserDef struct {
	Versionfile              int32
	Subversionfile           int32
	Flag                     int32
	Dupflag                  int32
	Savetime                 int32
	Tempdir                  [768]int8
	Fontdir                  [768]int8
	Renderdir                [1024]int8
	Textudir                 [768]int8
	Pythondir                [768]int8
	Sounddir                 [768]int8
	I18ndir                  [768]int8
	Image_editor             [1024]int8
	Anim_player              [1024]int8
	Anim_player_preset       int32
	V2d_min_gridsize         int16
	Timecode_style           int16
	Versions                 int16
	Dbl_click_time           int16
	Gameflags                int16
	Wheellinescroll          int16
	Uiflag                   int32
	Uiflag2                  int32
	Language                 int32
	Userpref                 int16
	Viewzoom                 int16
	Mixbufsize               int32
	Audiodevice              int32
	Audiorate                int32
	Audioformat              int32
	Audiochannels            int32
	Scrollback               int32
	Dpi                      int32
	Encoding                 int16
	Transopts                int16
	Menuthreshold1           int16
	Menuthreshold2           int16
	Themes                   ListBase
	Uifonts                  ListBase
	Uistyles                 ListBase
	Keymaps                  ListBase
	User_keymaps             ListBase
	Addons                   ListBase
	Keyconfigstr             [64]int8
	Undosteps                int16
	Undomemory               int16
	Gp_manhattendist         int16
	Gp_euclideandist         int16
	Gp_eraser                int16
	Gp_settings              int16
	Tb_leftmouse             int16
	Tb_rightmouse            int16
	Light                    [3]SolidLight
	Tw_hotspot               int16
	Tw_flag                  int16
	Tw_handlesize            int16
	Tw_size                  int16
	Textimeout               int16
	Texcollectrate           int16
	Wmdrawmethod             int16
	Dragthreshold            int16
	Memcachelimit            int32
	Prefetchframes           int32
	Frameserverport          int16
	Pad_rot_angle            int16
	Obcenter_dia             int16
	Rvisize                  int16
	Rvibright                int16
	Recent_files             int16
	Smooth_viewtx            int16
	Glreslimit               int16
	Curssize                 int16
	Color_picker_type        int16
	Ipo_new                  int16
	Keyhandles_new           int16
	Scrcastfps               int16
	Scrcastwait              int16
	Widget_unit              int16
	Anisotropic_filter       int16
	Use_16bit_textures       int16
	Use_gpu_mipmap           int16
	Ndof_sensitivity         float32
	Ndof_orbit_sensitivity   float32
	Ndof_flag                int32
	Ogl_multisamples         int16
	Image_draw_method        int16
	Glalphaclip              float32
	Autokey_mode             int16
	Autokey_flag             int16
	Text_render              int16
	Pad9                     int16
	Coba_weight              ColorBand
	Sculpt_paint_overlay_col [3]float32
	Tweak_threshold          int16
	Pad3                     int16
	Author                   [80]int8
	Compute_device_type      int32
	Compute_device_id        int32
	Fcu_inactive_alpha       float32
	Pixelsize                float32
}

// SDNA index: 220
type BScreen struct {
	Id                  ID
	Vertbase            ListBase
	Edgebase            ListBase
	Areabase            ListBase
	Regionbase          ListBase
	Scene               Pointer // *Scene
	Newscene            Pointer // *Scene
	Redraws_flag        int32
	Pad1                int32
	Full                int16
	Temp                int16
	Winid               int16
	Do_draw             int16
	Do_refresh          int16
	Do_draw_gesture     int16
	Do_draw_paintcursor int16
	Do_draw_drag        int16
	Swap                int16
	Mainwin             int16
	Subwinactive        int16
	Pad                 int16
	Animtimer           Pointer // *WmTimer
	Context             Pointer // *struct{}
}

// SDNA index: 221
type ScrVert struct {
	Next     Pointer // *ScrVert
	Prev     Pointer // *ScrVert
	Newv     Pointer // *ScrVert
	Vec      Vec2s
	Flag     int16
	Editflag int16
}

// SDNA index: 222
type ScrEdge struct {
	Next   Pointer // *ScrEdge
	Prev   Pointer // *ScrEdge
	V1     Pointer // *ScrVert
	V2     Pointer // *ScrVert
	Border int16
	Flag   int16
	Pad    int32
}

// SDNA index: 223
type Panel struct {
	Next         Pointer // *Panel
	Prev         Pointer // *Panel
	Type         Pointer // *PanelType
	Layout       Pointer // *UiLayout
	Panelname    [64]int8
	Tabname      [64]int8
	Drawname     [64]int8
	Ofsx         int32
	Ofsy         int32
	Sizex        int32
	Sizey        int32
	Labelofs     int16
	Pad          int16
	Flag         int16
	Runtime_flag int16
	Control      int16
	Snap         int16
	Sortorder    int32
	Paneltab     Pointer // *Panel
	Activedata   Pointer // *struct{}
}

// SDNA index: 224
type UiList struct {
	Next           Pointer // *UiList
	Prev           Pointer // *UiList
	Type           Pointer // *UiListType
	Padp           Pointer // *struct{}
	List_id        [64]int8
	Layout_type    int32
	Padi           int32
	List_scroll    int32
	List_size      int32
	List_last_len  int32
	List_grip_size int32
}

// SDNA index: 225
type ScrArea struct {
	Next              Pointer // *ScrArea
	Prev              Pointer // *ScrArea
	V1                Pointer // *ScrVert
	V2                Pointer // *ScrVert
	V3                Pointer // *ScrVert
	V4                Pointer // *ScrVert
	Full              Pointer // *BScreen
	Totrct            Rcti
	Spacetype         int8
	Butspacetype      int8
	Winx              int16
	Winy              int16
	Headertype        int16
	Do_refresh        int16
	Flag              int16
	Region_active_win int16
	Pad               int16
	Type              Pointer // *SpaceType
	Spacedata         ListBase
	Regionbase        ListBase
	Handlers          ListBase
	Actionzones       ListBase
}

// SDNA index: 226
type ARegion struct {
	Next            Pointer // *ARegion
	Prev            Pointer // *ARegion
	V2d             View2D
	Winrct          Rcti
	Drawrct         Rcti
	Winx            int16
	Winy            int16
	Swinid          int16
	Regiontype      int16
	Alignment       int16
	Flag            int16
	Fsize           float32
	Sizex           int16
	Sizey           int16
	Do_draw         int16
	Do_draw_overlay int16
	Swap            int16
	Overlap         int16
	Pad             [2]int16
	Type            Pointer // *ARegionType
	Uiblocks        ListBase
	Panels          ListBase
	Ui_lists        ListBase
	Handlers        ListBase
	Regiontimer     Pointer // *WmTimer
	Headerstr       Pointer // *int8
	Regiondata      Pointer // *struct{}
}

// SDNA index: 227
type FileGlobal struct {
	Subvstr       [4]int8
	Subversion    int16
	Pads          int16
	Minversion    int16
	Minsubversion int16
	Displaymode   int16
	Winpos        int16
	Curscreen     Pointer // *BScreen
	Curscene      Pointer // *Scene
	Fileflags     int32
	Globalf       int32
	Revision      int32
	Pad           int32
	Filename      [1024]int8
}

// SDNA index: 228
type StripElem struct {
	Name        [256]int8
	Orig_width  int32
	Orig_height int32
}

// SDNA index: 229
type StripCrop struct {
	Top    int32
	Bottom int32
	Left   int32
	Right  int32
}

// SDNA index: 230
type StripTransform struct {
	Xofs int32
	Yofs int32
}

// SDNA index: 231
type StripColorBalance struct {
	Lift  [3]float32
	Gamma [3]float32
	Gain  [3]float32
	Flag  int32
	Pad   int32
}

// SDNA index: 232
type StripProxy struct {
	Dir              [768]int8
	File             [256]int8
	Anim             Pointer // *Anim
	Tc               int16
	Quality          int16
	Build_size_flags int16
	Build_tc_flags   int16
}

// SDNA index: 233
type Strip struct {
	Next                Pointer // *Strip
	Prev                Pointer // *Strip
	Us                  int32
	Done                int32
	Startstill          int32
	Endstill            int32
	Stripdata           Pointer // *StripElem
	Dir                 [768]int8
	Proxy               Pointer // *StripProxy
	Crop                Pointer // *StripCrop
	Transform           Pointer // *StripTransform
	Color_balance       Pointer // *StripColorBalance
	Colorspace_settings ColorManagedColorspaceSettings
}

// SDNA index: 234
type Sequence struct {
	Next            Pointer // *Sequence
	Prev            Pointer // *Sequence
	Tmp             Pointer // *struct{}
	Lib             Pointer // *struct{}
	Name            [64]int8
	Flag            int32
	Type            int32
	Len             int32
	Start           int32
	Startofs        int32
	Endofs          int32
	Startstill      int32
	Endstill        int32
	Machine         int32
	Depth           int32
	Startdisp       int32
	Enddisp         int32
	Sat             float32
	Mul             float32
	Handsize        float32
	Anim_preseek    int16
	Streamindex     int16
	Multicam_source int32
	Clip_flag       int32
	Strip           Pointer // *Strip
	Ipo             Pointer // *Ipo
	Scene           Pointer // *Scene
	Scene_camera    Pointer // *Object
	Clip            Pointer // *MovieClip
	Mask            Pointer // *Mask
	Anim            Pointer // *Anim
	Effect_fader    float32
	Speed_fader     float32
	Seq1            Pointer // *Sequence
	Seq2            Pointer // *Sequence
	Seq3            Pointer // *Sequence
	Seqbase         ListBase
	Sound           Pointer // *BSound
	Scene_sound     Pointer // *struct{}
	Volume          float32
	Pitch           float32
	Pan             float32
	Strobe          float32
	Effectdata      Pointer // *struct{}
	Anim_startofs   int32
	Anim_endofs     int32
	Blend_mode      int32
	Blend_opacity   float32
	Sfra            int32
	Alpha_mode      int8
	Pad             [3]int8
	Modifiers       ListBase
}

// SDNA index: 235
type MetaStack struct {
	Next     Pointer // *MetaStack
	Prev     Pointer // *MetaStack
	Oldbasep Pointer // *ListBase
	Parseq   Pointer // *Sequence
}

// SDNA index: 236
type Editing struct {
	Seqbasep     Pointer // *ListBase
	Seqbase      ListBase
	Metastack    ListBase
	Act_seq      Pointer // *Sequence
	Act_imagedir [1024]int8
	Act_sounddir [1024]int8
	Over_ofs     int32
	Over_cfra    int32
	Over_flag    int32
	Pad          int32
	Over_border  Rctf
}

// SDNA index: 237
type WipeVars struct {
	EdgeWidth float32
	Angle     float32
	Forward   int16
	Wipetype  int16
}

// SDNA index: 238
type GlowVars struct {
	FMini    float32
	FClamp   float32
	FBoost   float32
	DDist    float32
	DQuality int32
	BNoComp  int32
}

// SDNA index: 239
type TransformVars struct {
	ScalexIni     float32
	ScaleyIni     float32
	XIni          float32
	YIni          float32
	RotIni        float32
	Percent       int32
	Interpolation int32
	Uniform_scale int32
}

// SDNA index: 240
type SolidColorVars struct {
	Col [3]float32
	Pad float32
}

// SDNA index: 241
type SpeedControlVars struct {
	FrameMap       Pointer // *float32
	GlobalSpeed    float32
	Flags          int32
	Length         int32
	LastValidFrame int32
}

// SDNA index: 242
type SequenceModifierData struct {
	Next            Pointer // *SequenceModifierData
	Prev            Pointer // *SequenceModifierData
	Type            int32
	Flag            int32
	Name            [64]int8
	Mask_input_type int32
	Pad             int32
	Mask_sequence   Pointer // *Sequence
	Mask_id         Pointer // *Mask
}

// SDNA index: 243
type ColorBalanceModifierData struct {
	Modifier       SequenceModifierData
	Color_balance  StripColorBalance
	Color_multiply float32
}

// SDNA index: 244
type CurvesModifierData struct {
	Modifier      SequenceModifierData
	Curve_mapping CurveMapping
}

// SDNA index: 245
type HueCorrectModifierData struct {
	Modifier      SequenceModifierData
	Curve_mapping CurveMapping
}

// SDNA index: 246
type BrightContrastModifierData struct {
	Modifier SequenceModifierData
	Bright   float32
	Contrast float32
}

// SDNA index: 247
type SequencerScopes struct {
	Reference_ibuf    Pointer // *ImBuf
	Zebra_ibuf        Pointer // *ImBuf
	Waveform_ibuf     Pointer // *ImBuf
	Sep_waveform_ibuf Pointer // *ImBuf
	Vector_ibuf       Pointer // *ImBuf
	Histogram_ibuf    Pointer // *ImBuf
}

// SDNA index: 248
type Effect struct {
	Next    Pointer // *Effect
	Prev    Pointer // *Effect
	Type    int16
	Flag    int16
	Buttype int16
	Rt      int16
}

// SDNA index: 249
type BuildEff struct {
	Next    Pointer // *BuildEff
	Prev    Pointer // *BuildEff
	Type    int16
	Flag    int16
	Buttype int16
	Rt      int16
	Len     float32
	Sfra    float32
}

// SDNA index: 250
type PartEff struct {
	Next         Pointer // *PartEff
	Prev         Pointer // *PartEff
	Type         int16
	Flag         int16
	Buttype      int16
	Stype        int16
	Vertgroup    int16
	Userjit      int16
	Sta          float32
	End          float32
	Lifetime     float32
	Totpart      int32
	Totkey       int32
	Seed         int32
	Normfac      float32
	Obfac        float32
	Randfac      float32
	Texfac       float32
	Randlife     float32
	Force        [3]float32
	Damp         float32
	Nabla        float32
	Vectsize     float32
	Maxlen       float32
	Pad          float32
	Defvec       [3]float32
	Mult         [4]float32
	Life         [4]float32
	Child        [4]int16
	Mat          [4]int16
	Texmap       int16
	Curmult      int16
	Staticstep   int16
	Omat         int16
	Timetex      int16
	Speedtex     int16
	Flag2        int16
	Flag2neg     int16
	Disp         int16
	Vertgroup_v  int16
	Vgroupname   [64]int8
	Vgroupname_v [64]int8
	Imat         [4][4]float32
	Keys         Pointer // *Particle
	Group        Pointer // *Group
}

// SDNA index: 251
type WaveEff struct {
	Next     Pointer // *WaveEff
	Prev     Pointer // *WaveEff
	Type     int16
	Flag     int16
	Buttype  int16
	Stype    int16
	Startx   float32
	Starty   float32
	Height   float32
	Width    float32
	Narrow   float32
	Speed    float32
	Minfac   float32
	Damp     float32
	Timeoffs float32
	Lifetime float32
}

// SDNA index: 252
type TreeStoreElem struct {
	Type int16
	Nr   int16
	Flag int16
	Used int16
	Id   Pointer // *ID
}

// SDNA index: 253
type TreeStore struct {
	Totelem  int32
	Usedelem int32
	Data     Pointer // *TreeStoreElem
}

// SDNA index: 254
type BProperty struct {
	Next Pointer // *BProperty
	Prev Pointer // *BProperty
	Name [64]int8
	Type int16
	Flag int16
	Data int32
	Poin Pointer // *struct{}
}

// SDNA index: 255
type BNearSensor struct {
	Name      [64]int8
	Dist      float32
	Resetdist float32
	Lastval   int32
	Pad       int32
}

// SDNA index: 256
type BMouseSensor struct {
	Type int16
	Flag int16
	Pad1 int16
	Pad2 int16
}

// SDNA index: 257
type BTouchSensor struct {
	Name [64]int8
	Ma   Pointer // *Material
	Dist float32
	Pad  float32
}

// SDNA index: 258
type BKeyboardSensor struct {
	Key        int16
	Qual       int16
	Type       int16
	Qual2      int16
	TargetName [64]int8
	ToggleName [64]int8
}

// SDNA index: 259
type BPropertySensor struct {
	Type     int32
	Pad      int32
	Name     [64]int8
	Value    [64]int8
	Maxvalue [64]int8
}

// SDNA index: 260
type BActuatorSensor struct {
	Type int32
	Pad  int32
	Name [64]int8
}

// SDNA index: 261
type BDelaySensor struct {
	Delay    int16
	Duration int16
	Flag     int16
	Pad      int16
}

// SDNA index: 262
type BCollisionSensor struct {
	Name         [64]int8
	MaterialName [64]int8
	Damptimer    int16
	Damp         int16
	Mode         int16
	Pad2         int16
}

// SDNA index: 263
type BRadarSensor struct {
	Name  [64]int8
	Angle float32
	Range float32
	Flag  int16
	Axis  int16
}

// SDNA index: 264
type BRandomSensor struct {
	Name  [64]int8
	Seed  int32
	Delay int32
}

// SDNA index: 265
type BRaySensor struct {
	Name     [64]int8
	Range    float32
	Propname [64]int8
	Matname  [64]int8
	Mode     int16
	Pad1     int16
	Axisflag int32
}

// SDNA index: 266
type BArmatureSensor struct {
	Posechannel [64]int8
	Constraint  [64]int8
	Type        int32
	Value       float32
}

// SDNA index: 267
type BMessageSensor struct {
	FromObject Pointer // *Object
	Subject    [64]int8
	Body       [64]int8
}

// SDNA index: 268
type BSensor struct {
	Next     Pointer // *BSensor
	Prev     Pointer // *BSensor
	Type     int16
	Otype    int16
	Flag     int16
	Pulse    int16
	Freq     int16
	Totlinks int16
	Pad1     int16
	Pad2     int16
	Name     [64]int8
	Data     Pointer // *struct{}
	Links    Pointer // **BController
	Ob       Pointer // *Object
	Invert   int16
	Level    int16
	Tap      int16
	Pad      int16
}

// SDNA index: 269
type BJoystickSensor struct {
	Name        [64]int8
	Type        int8
	Joyindex    int8
	Flag        int16
	Axis        int16
	Axis_single int16
	Axisf       int32
	Button      int32
	Hat         int32
	Hatf        int32
	Precision   int32
}

// SDNA index: 270
type BExpressionCont struct {
	Str [128]int8
}

// SDNA index: 271
type BPythonCont struct {
	Text   Pointer // *Text
	Module [64]int8
	Mode   int32
	Flag   int32
}

// SDNA index: 272
type BController struct {
	Next       Pointer // *BController
	Prev       Pointer // *BController
	Mynew      Pointer // *BController
	Type       int16
	Flag       int16
	Inputs     int16
	Totlinks   int16
	Otype      int16
	Totslinks  int16
	Pad2       int16
	Pad3       int16
	Name       [64]int8
	Data       Pointer // *struct{}
	Links      Pointer // **BActuator
	Slinks     Pointer // **BSensor
	Val        int16
	Valo       int16
	State_mask int32
}

// SDNA index: 273
type BAddObjectActuator struct {
	Time int32
	Pad  int32
	Ob   Pointer // *Object
}

// SDNA index: 274
type BActionActuator struct {
	Act          Pointer // *BAction
	Type         int16
	Flag         int16
	Sta          float32
	End          float32
	Name         [64]int8
	FrameProp    [64]int8
	Blendin      int16
	Priority     int16
	Layer        int16
	End_reset    int16
	Strideaxis   int16
	Pad          int16
	Stridelength float32
	Layer_weight float32
}

// SDNA index: 275
type Sound3D struct {
	Min_gain           float32
	Max_gain           float32
	Reference_distance float32
	Max_distance       float32
	Rolloff_factor     float32
	Cone_inner_angle   float32
	Cone_outer_angle   float32
	Cone_outer_gain    float32
}

// SDNA index: 276
type BSoundActuator struct {
	Flag    int16
	Sndnr   int16
	Pad1    int32
	Pad2    int32
	Pad3    [2]int16
	Volume  float32
	Pitch   float32
	Sound   Pointer // *BSound
	Sound3D Sound3D
	Type    int16
	Pad4    int16
	Pad5    int16
	Pad6    [1]int16
}

// SDNA index: 277
type BEditObjectActuator struct {
	Time          int32
	Type          int16
	Flag          int16
	Ob            Pointer // *Object
	Me            Pointer // *Mesh
	Name          [64]int8
	LinVelocity   [3]float32
	AngVelocity   [3]float32
	Mass          float32
	Localflag     int16
	Dyn_operation int16
}

// SDNA index: 278
type BSceneActuator struct {
	Type   int16
	Pad1   int16
	Pad    int32
	Scene  Pointer // *Scene
	Camera Pointer // *Object
}

// SDNA index: 279
type BPropertyActuator struct {
	Pad   int32
	Type  int32
	Name  [64]int8
	Value [64]int8
	Ob    Pointer // *Object
}

// SDNA index: 280
type BObjectActuator struct {
	Flag            int16
	Type            int16
	Otype           int16
	Damping         int16
	Forceloc        [3]float32
	Forcerot        [3]float32
	Pad             [3]float32
	Pad1            [3]float32
	Dloc            [3]float32
	Drot            [3]float32
	Linearvelocity  [3]float32
	Angularvelocity [3]float32
	Reference       Pointer // *Object
}

// SDNA index: 281
type BIpoActuator struct {
	Flag      int16
	Type      int16
	Sta       float32
	End       float32
	Name      [64]int8
	FrameProp [64]int8
	Pad1      int16
	Pad2      int16
	Pad3      int16
	Pad4      int16
}

// SDNA index: 282
type BCameraActuator struct {
	Ob      Pointer // *Object
	Height  float32
	Min     float32
	Max     float32
	Damping float32
	Pad1    int16
	Axis    int16
	Pad2    float32
}

// SDNA index: 283
type BConstraintActuator struct {
	Type    int16
	Mode    int16
	Flag    int16
	Damp    int16
	Time    int16
	Rotdamp int16
	Pad     int32
	Minloc  [3]float32
	Maxloc  [3]float32
	Minrot  [3]float32
	Maxrot  [3]float32
	Matprop [64]int8
}

// SDNA index: 284
type BGroupActuator struct {
	Flag   int16
	Type   int16
	Sta    int32
	End    int32
	Name   [64]int8
	Pad    [3]int16
	Cur    int16
	Butsta int16
	Butend int16
}

// SDNA index: 285
type BRandomActuator struct {
	Seed         int32
	Distribution int32
	Int_arg_1    int32
	Int_arg_2    int32
	Float_arg_1  float32
	Float_arg_2  float32
	Propname     [64]int8
}

// SDNA index: 286
type BMessageActuator struct {
	ToPropName [64]int8
	ToObject   Pointer // *Object
	Subject    [64]int8
	BodyType   int16
	Pad1       int16
	Pad2       int32
	Body       [64]int8
}

// SDNA index: 287
type BGameActuator struct {
	Flag        int16
	Type        int16
	Sta         int32
	End         int32
	Filename    [64]int8
	Loadaniname [64]int8
}

// SDNA index: 288
type BVisibilityActuator struct {
	Flag int32
}

// SDNA index: 289
type BTwoDFilterActuator struct {
	Pad       [4]int8
	Type      int16
	Flag      int16
	Int_arg   int32
	Float_arg float32
	Text      Pointer // *Text
}

// SDNA index: 290
type BParentActuator struct {
	Pad  [2]int8
	Flag int16
	Type int32
	Ob   Pointer // *Object
}

// SDNA index: 291
type BStateActuator struct {
	Type int32
	Mask int32
}

// SDNA index: 292
type BArmatureActuator struct {
	Posechannel [64]int8
	Constraint  [64]int8
	Type        int32
	Weight      float32
	Influence   float32
	Pad         float32
	Target      Pointer // *Object
	Subtarget   Pointer // *Object
}

// SDNA index: 293
type BSteeringActuator struct {
	Pad          [5]int8
	Flag         int8
	Facingaxis   int16
	Type         int32
	Dist         float32
	Velocity     float32
	Acceleration float32
	Turnspeed    float32
	UpdateTime   int32
	Target       Pointer // *Object
	Navmesh      Pointer // *Object
}

// SDNA index: 294
type BActuator struct {
	Next  Pointer // *BActuator
	Prev  Pointer // *BActuator
	Mynew Pointer // *BActuator
	Type  int16
	Flag  int16
	Otype int16
	Go    int16
	Name  [64]int8
	Data  Pointer // *struct{}
	Ob    Pointer // *Object
}

// SDNA index: 295
type BSound struct {
	Id              ID
	Name            [1024]int8
	Packedfile      Pointer // *PackedFile
	Handle          Pointer // *struct{}
	Newpackedfile   Pointer // *PackedFile
	Ipo             Pointer // *Ipo
	Volume          float32
	Attenuation     float32
	Pitch           float32
	Min_gain        float32
	Max_gain        float32
	Distance        float32
	Flags           int32
	Pad             int32
	Cache           Pointer // *struct{}
	Waveform        Pointer // *struct{}
	Playback_handle Pointer // *struct{}
}

// SDNA index: 296
type GroupObject struct {
	Next    Pointer // *GroupObject
	Prev    Pointer // *GroupObject
	Ob      Pointer // *Object
	Lampren Pointer // *struct{}
	Recalc  int16
	Pad     [6]int8
}

// SDNA index: 297
type Group struct {
	Id        ID
	Gobject   ListBase
	Layer     int32
	Dupli_ofs [3]float32
}

// SDNA index: 298
type Bone struct {
	Next      Pointer // *Bone
	Prev      Pointer // *Bone
	Prop      Pointer // *IDProperty
	Parent    Pointer // *Bone
	Childbase ListBase
	Name      [64]int8
	Roll      float32
	Head      [3]float32
	Tail      [3]float32
	Bone_mat  [3][3]float32
	Flag      int32
	Arm_head  [3]float32
	Arm_tail  [3]float32
	Arm_mat   [4][4]float32
	Arm_roll  float32
	Dist      float32
	Weight    float32
	Xwidth    float32
	Length    float32
	Zwidth    float32
	Ease1     float32
	Ease2     float32
	Rad_head  float32
	Rad_tail  float32
	Size      [3]float32
	Layer     int32
	Segments  int16
	Pad       [1]int16
}

// SDNA index: 299
type BArmature struct {
	Id              ID
	Adt             Pointer // *AnimData
	Bonebase        ListBase
	Chainbase       ListBase
	Edbo            Pointer // *ListBase
	Act_bone        Pointer // *Bone
	Act_edbone      Pointer // *struct{}
	Sketch          Pointer // *struct{}
	Flag            int32
	Drawtype        int32
	Gevertdeformer  int32
	Pad             int32
	Deformflag      int16
	Pathflag        int16
	Layer_used      int32
	Layer           int32
	Layer_protected int32
	Ghostep         int16
	Ghostsize       int16
	Ghosttype       int16
	Pathsize        int16
	Ghostsf         int32
	Ghostef         int32
	Pathsf          int32
	Pathef          int32
	Pathbc          int32
	Pathac          int32
}

// SDNA index: 300
type BMotionPathVert struct {
	Co   [3]float32
	Flag int32
}

// SDNA index: 301
type BMotionPath struct {
	Points      Pointer // *BMotionPathVert
	Length      int32
	Start_frame int32
	End_frame   int32
	Flag        int32
}

// SDNA index: 302
type BAnimVizSettings struct {
	Ghost_sf      int32
	Ghost_ef      int32
	Ghost_bc      int32
	Ghost_ac      int32
	Ghost_type    int16
	Ghost_step    int16
	Ghost_flag    int16
	Recalc        int16
	Path_type     int16
	Path_step     int16
	Path_viewflag int16
	Path_bakeflag int16
	Path_sf       int32
	Path_ef       int32
	Path_bc       int32
	Path_ac       int32
}

// SDNA index: 303
type BPoseChannel struct {
	Next        Pointer // *BPoseChannel
	Prev        Pointer // *BPoseChannel
	Prop        Pointer // *IDProperty
	Constraints ListBase
	Name        [64]int8
	Flag        int16
	Ikflag      int16
	Protectflag int16
	Agrp_index  int16
	Constflag   int8
	Selectflag  int8
	Pad0        [6]int8
	Bone        Pointer // *Bone
	Parent      Pointer // *BPoseChannel
	Child       Pointer // *BPoseChannel
	Iktree      ListBase
	Siktree     ListBase
	Mpath       Pointer // *BMotionPath
	Custom      Pointer // *Object
	Custom_tx   Pointer // *BPoseChannel
	Loc         [3]float32
	Size        [3]float32
	Eul         [3]float32
	Quat        [4]float32
	RotAxis     [3]float32
	RotAngle    float32
	Rotmode     int16
	Pad         int16
	Chan_mat    [4][4]float32
	Pose_mat    [4][4]float32
	Constinv    [4][4]float32
	Pose_head   [3]float32
	Pose_tail   [3]float32
	Limitmin    [3]float32
	Limitmax    [3]float32
	Stiffness   [3]float32
	Ikstretch   float32
	Ikrotweight float32
	Iklinweight float32
	Temp        Pointer // *struct{}
}

// SDNA index: 304
type BPose struct {
	Chanbase       ListBase
	Chanhash       Pointer // *GHash
	Flag           int16
	Pad            int16
	Proxy_layer    int32
	Pad1           int32
	Ctime          float32
	Stride_offset  [3]float32
	Cyclic_offset  [3]float32
	Agroups        ListBase
	Active_group   int32
	Iksolver       int32
	Ikdata         Pointer // *struct{}
	Ikparam        Pointer // *struct{}
	Avs            BAnimVizSettings
	Proxy_act_bone [64]int8
}

// SDNA index: 305
type BIKParam struct {
	Iksolver int32
}

// SDNA index: 306
type BItasc struct {
	Iksolver  int32
	Precision float32
	Numiter   int16
	Numstep   int16
	Minstep   float32
	Maxstep   float32
	Solver    int16
	Flag      int16
	Feedback  float32
	Maxvel    float32
	Dampmax   float32
	Dampeps   float32
}

// SDNA index: 307
type BActionGroup struct {
	Next      Pointer // *BActionGroup
	Prev      Pointer // *BActionGroup
	Channels  ListBase
	Flag      int32
	CustomCol int32
	Name      [64]int8
	Cs        ThemeWireColor
}

// SDNA index: 308
type BAction struct {
	Id            ID
	Curves        ListBase
	Chanbase      ListBase
	Groups        ListBase
	Markers       ListBase
	Flag          int32
	Active_marker int32
	Idroot        int32
	Pad           int32
}

// SDNA index: 309
type BDopeSheet struct {
	Source      Pointer // *ID
	Chanbase    ListBase
	Filter_grp  Pointer // *Group
	Searchstr   [64]int8
	Filterflag  int32
	Flag        int32
	RenameIndex int32
	Pad         int32
}

// SDNA index: 310
type SpaceAction struct {
	Next         Pointer // *SpaceLink
	Prev         Pointer // *SpaceLink
	Regionbase   ListBase
	Spacetype    int32
	Blockscale   float32
	Blockhandler [8]int16
	V2d          View2D
	Action       Pointer // *BAction
	Ads          BDopeSheet
	Mode         int8
	Autosnap     int8
	Flag         int16
	Timeslide    float32
}

// SDNA index: 311
type BActionChannel struct {
	Next               Pointer // *BActionChannel
	Prev               Pointer // *BActionChannel
	Grp                Pointer // *BActionGroup
	Ipo                Pointer // *Ipo
	ConstraintChannels ListBase
	Flag               int32
	Name               [64]int8
	Temp               int32
}

// SDNA index: 312
type BConstraintChannel struct {
	Next Pointer // *BConstraintChannel
	Prev Pointer // *BConstraintChannel
	Ipo  Pointer // *Ipo
	Flag int16
	Name [30]int8
}

// SDNA index: 313
type BConstraint struct {
	Next      Pointer // *BConstraint
	Prev      Pointer // *BConstraint
	Data      Pointer // *struct{}
	Type      int16
	Flag      int16
	Ownspace  int8
	Tarspace  int8
	Name      [64]int8
	Pad       int16
	Enforce   float32
	Headtail  float32
	Ipo       Pointer // *Ipo
	Lin_error float32
	Rot_error float32
}

// SDNA index: 314
type BConstraintTarget struct {
	Next      Pointer // *BConstraintTarget
	Prev      Pointer // *BConstraintTarget
	Tar       Pointer // *Object
	Subtarget [64]int8
	Matrix    [4][4]float32
	Space     int16
	Flag      int16
	Type      int16
	RotOrder  int16
}

// SDNA index: 315
type BPythonConstraint struct {
	Text      Pointer // *Text
	Prop      Pointer // *IDProperty
	Flag      int32
	Tarnum    int32
	Targets   ListBase
	Tar       Pointer // *Object
	Subtarget [64]int8
}

// SDNA index: 316
type BKinematicConstraint struct {
	Tar           Pointer // *Object
	Iterations    int16
	Flag          int16
	Rootbone      int16
	Max_rootbone  int16
	Subtarget     [64]int8
	Poletar       Pointer // *Object
	Polesubtarget [64]int8
	Poleangle     float32
	Weight        float32
	Orientweight  float32
	Grabtarget    [3]float32
	Type          int16
	Mode          int16
	Dist          float32
}

// SDNA index: 317
type BSplineIKConstraint struct {
	Tar         Pointer // *Object
	Points      Pointer // *float32
	Numpoints   int16
	Chainlen    int16
	Flag        int16
	XzScaleMode int16
}

// SDNA index: 318
type BTrackToConstraint struct {
	Tar       Pointer // *Object
	Reserved1 int32
	Reserved2 int32
	Flags     int32
	Pad       int32
	Subtarget [64]int8
}

// SDNA index: 319
type BRotateLikeConstraint struct {
	Tar       Pointer // *Object
	Flag      int32
	Reserved1 int32
	Subtarget [64]int8
}

// SDNA index: 320
type BLocateLikeConstraint struct {
	Tar       Pointer // *Object
	Flag      int32
	Reserved1 int32
	Subtarget [64]int8
}

// SDNA index: 321
type BSizeLikeConstraint struct {
	Tar       Pointer // *Object
	Flag      int32
	Reserved1 int32
	Subtarget [64]int8
}

// SDNA index: 322
type BSameVolumeConstraint struct {
	Flag   int32
	Volume float32
}

// SDNA index: 323
type BTransLikeConstraint struct {
	Tar       Pointer // *Object
	Subtarget [64]int8
}

// SDNA index: 324
type BMinMaxConstraint struct {
	Tar        Pointer // *Object
	Minmaxflag int32
	Offset     float32
	Flag       int32
	Sticky     int16
	Stuck      int16
	Pad1       int16
	Pad2       int16
	Cache      [3]float32
	Subtarget  [64]int8
}

// SDNA index: 325
type BActionConstraint struct {
	Tar       Pointer // *Object
	Type      int16
	Local     int16
	Start     int32
	End       int32
	Min       float32
	Max       float32
	Flag      int32
	Act       Pointer // *BAction
	Subtarget [64]int8
}

// SDNA index: 326
type BLockTrackConstraint struct {
	Tar       Pointer // *Object
	Trackflag int32
	Lockflag  int32
	Subtarget [64]int8
}

// SDNA index: 327
type BDampTrackConstraint struct {
	Tar       Pointer // *Object
	Trackflag int32
	Pad       int32
	Subtarget [64]int8
}

// SDNA index: 328
type BFollowPathConstraint struct {
	Tar        Pointer // *Object
	Offset     float32
	Offset_fac float32
	Followflag int32
	Trackflag  int16
	Upflag     int16
}

// SDNA index: 329
type BStretchToConstraint struct {
	Tar       Pointer // *Object
	Volmode   int32
	Plane     int32
	Orglength float32
	Bulge     float32
	Subtarget [64]int8
}

// SDNA index: 330
type BRigidBodyJointConstraint struct {
	Tar      Pointer // *Object
	Child    Pointer // *Object
	Type     int32
	PivX     float32
	PivY     float32
	PivZ     float32
	AxX      float32
	AxY      float32
	AxZ      float32
	MinLimit [6]float32
	MaxLimit [6]float32
	ExtraFz  float32
	Flag     int16
	Pad      int16
	Pad1     int16
	Pad2     int16
}

// SDNA index: 331
type BClampToConstraint struct {
	Tar   Pointer // *Object
	Flag  int32
	Flag2 int32
}

// SDNA index: 332
type BChildOfConstraint struct {
	Tar       Pointer // *Object
	Flag      int32
	Pad       int32
	Invmat    [4][4]float32
	Subtarget [64]int8
}

// SDNA index: 333
type BTransformConstraint struct {
	Tar       Pointer // *Object
	Subtarget [64]int8
	From      int16
	To        int16
	Map       [3]int8
	Expo      int8
	From_min  [3]float32
	From_max  [3]float32
	To_min    [3]float32
	To_max    [3]float32
}

// SDNA index: 334
type BPivotConstraint struct {
	Tar       Pointer // *Object
	Subtarget [64]int8
	Offset    [3]float32
	RotAxis   int16
	Flag      int16
}

// SDNA index: 335
type BLocLimitConstraint struct {
	Xmin  float32
	Xmax  float32
	Ymin  float32
	Ymax  float32
	Zmin  float32
	Zmax  float32
	Flag  int16
	Flag2 int16
}

// SDNA index: 336
type BRotLimitConstraint struct {
	Xmin  float32
	Xmax  float32
	Ymin  float32
	Ymax  float32
	Zmin  float32
	Zmax  float32
	Flag  int16
	Flag2 int16
}

// SDNA index: 337
type BSizeLimitConstraint struct {
	Xmin  float32
	Xmax  float32
	Ymin  float32
	Ymax  float32
	Zmin  float32
	Zmax  float32
	Flag  int16
	Flag2 int16
}

// SDNA index: 338
type BDistLimitConstraint struct {
	Tar       Pointer // *Object
	Subtarget [64]int8
	Dist      float32
	Soft      float32
	Flag      int16
	Mode      int16
	Pad       int32
}

// SDNA index: 339
type BShrinkwrapConstraint struct {
	Target     Pointer // *Object
	Dist       float32
	ShrinkType int16
	ProjAxis   int8
	Pad        [9]int8
}

// SDNA index: 340
type BFollowTrackConstraint struct {
	Clip         Pointer // *MovieClip
	Track        [64]int8
	Flag         int32
	Frame_method int32
	Object       [64]int8
	Camera       Pointer // *Object
	Depth_ob     Pointer // *Object
}

// SDNA index: 341
type BCameraSolverConstraint struct {
	Clip Pointer // *MovieClip
	Flag int32
	Pad  int32
}

// SDNA index: 342
type BObjectSolverConstraint struct {
	Clip   Pointer // *MovieClip
	Flag   int32
	Pad    int32
	Object [64]int8
	Invmat [4][4]float32
	Camera Pointer // *Object
}

// SDNA index: 343
type BActionModifier struct {
	Next        Pointer // *BActionModifier
	Prev        Pointer // *BActionModifier
	Type        int16
	Flag        int16
	Channel     [32]int8
	Noisesize   float32
	Turbul      float32
	Channels    int16
	No_rot_axis int16
	Ob          Pointer // *Object
}

// SDNA index: 344
type BActionStrip struct {
	Next          Pointer // *BActionStrip
	Prev          Pointer // *BActionStrip
	Flag          int16
	Mode          int16
	Stride_axis   int16
	Curmod        int16
	Ipo           Pointer // *Ipo
	Act           Pointer // *BAction
	Object        Pointer // *Object
	Start         float32
	End           float32
	Actstart      float32
	Actend        float32
	Actoffs       float32
	Stridelen     float32
	Repeat        float32
	Scale         float32
	Blendin       float32
	Blendout      float32
	Stridechannel [32]int8
	Offs_bone     [32]int8
	Modifiers     ListBase
}

// SDNA index: 345
type BNodeStack struct {
	Vec        [4]float32
	Min        float32
	Max        float32
	Data       Pointer // *struct{}
	Hasinput   int16
	Hasoutput  int16
	Datatype   int16
	Sockettype int16
	Is_copy    int16
	External   int16
	Pad        [2]int16
}

// SDNA index: 346
type BNodeSocket struct {
	Next          Pointer // *BNodeSocket
	Prev          Pointer // *BNodeSocket
	New_sock      Pointer // *BNodeSocket
	Prop          Pointer // *IDProperty
	Identifier    [64]int8
	Name          [64]int8
	Storage       Pointer // *struct{}
	Type          int16
	Flag          int16
	Limit         int16
	In_out        int16
	Typeinfo      Pointer // *BNodeSocketType
	Idname        [64]int8
	Locx          float32
	Locy          float32
	Default_value Pointer // *struct{}
	Stack_index   int16
	Stack_type    int16
	Resizemode    int32
	Cache         Pointer // *struct{}
	Own_index     int32
	To_index      int32
	Groupsock     Pointer // *BNodeSocket
	Link          Pointer // *BNodeLink
	Ns            BNodeStack
}

// SDNA index: 347
type BNode struct {
	Next           Pointer // *BNode
	Prev           Pointer // *BNode
	New_node       Pointer // *BNode
	Prop           Pointer // *IDProperty
	Typeinfo       Pointer // *BNodeType
	Idname         [64]int8
	Name           [64]int8
	Flag           int32
	Type           int16
	Pad            int16
	Done           int16
	Level          int16
	Lasty          int16
	Menunr         int16
	Stack_index    int16
	Nr             int16
	Color          [3]float32
	Inputs         ListBase
	Outputs        ListBase
	Parent         Pointer // *BNode
	Id             Pointer // *ID
	Storage        Pointer // *struct{}
	Original       Pointer // *BNode
	Internal_links ListBase
	Locx           float32
	Locy           float32
	Width          float32
	Height         float32
	Miniwidth      float32
	Offsetx        float32
	Offsety        float32
	Update         int32
	Label          [64]int8
	Custom1        int16
	Custom2        int16
	Custom3        float32
	Custom4        float32
	Need_exec      int16
	Exec           int16
	Threaddata     Pointer // *struct{}
	Totr           Rctf
	Butr           Rctf
	Prvr           Rctf
	Preview_xsize  int16
	Preview_ysize  int16
	Pad2           int32
	Block          Pointer // *UiBlock
}

// SDNA index: 348
type BNodeInstanceKey struct {
	Value int32
}

// SDNA index: 349
type BNodeInstanceHashEntry struct {
	Key BNodeInstanceKey
	Tag int16
	Pad int16
}

// SDNA index: 350
type BNodePreview struct {
	Hash_entry BNodeInstanceHashEntry
	Rect       Pointer // *int8
	Xsize      int16
	Ysize      int16
	Pad        int32
}

// SDNA index: 351
type BNodeLink struct {
	Next     Pointer // *BNodeLink
	Prev     Pointer // *BNodeLink
	Fromnode Pointer // *BNode
	Tonode   Pointer // *BNode
	Fromsock Pointer // *BNodeSocket
	Tosock   Pointer // *BNodeSocket
	Flag     int32
	Pad      int32
}

// SDNA index: 352
type BNodeTree struct {
	Id                ID
	Adt               Pointer // *AnimData
	Typeinfo          Pointer // *BNodeTreeType
	Idname            [64]int8
	Interface_type    Pointer // *StructRNA
	Gpd               Pointer // *BGPdata
	View_center       [2]float32
	Nodes             ListBase
	Links             ListBase
	Type              int32
	Init              int32
	Cur_index         int32
	Flag              int32
	Update            int32
	Is_updating       int16
	Done              int16
	Pad2              int32
	Nodetype          int32
	Edit_quality      int16
	Render_quality    int16
	Chunksize         int32
	Viewer_border     Rctf
	Inputs            ListBase
	Outputs           ListBase
	Previews          Pointer // *BNodeInstanceHash
	Active_viewer_key BNodeInstanceKey
	Pad               int32
	Execdata          Pointer // *BNodeTreeExec
	Progress          func()
	Stats_draw        func()
	Test_break        func() int32
	Update_draw       func()
	Tbh               Pointer // *struct{}
	Prh               Pointer // *struct{}
	Sdh               Pointer // *struct{}
	Udh               Pointer // *struct{}
}

// SDNA index: 353
type BNodeSocketValueInt struct {
	Subtype int32
	Value   int32
	Min     int32
	Max     int32
}

// SDNA index: 354
type BNodeSocketValueFloat struct {
	Subtype int32
	Value   float32
	Min     float32
	Max     float32
}

// SDNA index: 355
type BNodeSocketValueBoolean struct {
	Value int8
	Pad   [3]int8
}

// SDNA index: 356
type BNodeSocketValueVector struct {
	Subtype int32
	Value   [3]float32
	Min     float32
	Max     float32
}

// SDNA index: 357
type BNodeSocketValueRGBA struct {
	Value [4]float32
}

// SDNA index: 358
type BNodeSocketValueString struct {
	Subtype int32
	Pad     int32
	Value   [1024]int8
}

// SDNA index: 359
type NodeFrame struct {
	Flag       int16
	Label_size int16
}

// SDNA index: 360
type NodeImageAnim struct {
	Frames int32
	Sfra   int32
	Nr     int32
	Cyclic int8
	Movie  int8
	Pad    int16
}

// SDNA index: 361
type ColorCorrectionData struct {
	Saturation float32
	Contrast   float32
	Gamma      float32
	Gain       float32
	Lift       float32
	Pad        int32
}

// SDNA index: 362
type NodeColorCorrection struct {
	Master        ColorCorrectionData
	Shadows       ColorCorrectionData
	Midtones      ColorCorrectionData
	Highlights    ColorCorrectionData
	Startmidtones float32
	Endmidtones   float32
}

// SDNA index: 363
type NodeBokehImage struct {
	Angle        float32
	Flaps        int32
	Rounding     float32
	Catadioptric float32
	Lensshift    float32
}

// SDNA index: 364
type NodeBoxMask struct {
	X        float32
	Y        float32
	Rotation float32
	Height   float32
	Width    float32
	Pad      int32
}

// SDNA index: 365
type NodeEllipseMask struct {
	X        float32
	Y        float32
	Rotation float32
	Height   float32
	Width    float32
	Pad      int32
}

// SDNA index: 366
type NodeImageLayer struct {
	Pass_index int32
	Pass_flag  int32
}

// SDNA index: 367
type NodeBlurData struct {
	Sizex           int16
	Sizey           int16
	Samples         int16
	Maxspeed        int16
	Minspeed        int16
	Relative        int16
	Aspect          int16
	Curved          int16
	Fac             float32
	Percentx        float32
	Percenty        float32
	Filtertype      int16
	Bokeh           int8
	Gamma           int8
	Image_in_width  int32
	Image_in_height int32
}

// SDNA index: 368
type NodeDBlurData struct {
	Center_x float32
	Center_y float32
	Distance float32
	Angle    float32
	Spin     float32
	Zoom     float32
	Iter     int16
	Wrap     int8
	Pad      int8
}

// SDNA index: 369
type NodeBilateralBlurData struct {
	Sigma_color float32
	Sigma_space float32
	Iter        int16
	Pad         int16
}

// SDNA index: 370
type NodeHueSat struct {
	Hue float32
	Sat float32
	Val float32
}

// SDNA index: 371
type NodeImageFile struct {
	Name      [1024]int8
	Im_format ImageFormatData
	Sfra      int32
	Efra      int32
}

// SDNA index: 372
type NodeImageMultiFile struct {
	Base_path    [1024]int8
	Format       ImageFormatData
	Sfra         int32
	Efra         int32
	Active_input int32
	Pad          int32
}

// SDNA index: 373
type NodeImageMultiFileSocket struct {
	Use_render_format int16
	Use_node_format   int16
	Pad1              int32
	Path              [1024]int8
	Format            ImageFormatData
	Layer             [30]int8
	Pad2              [2]int8
}

// SDNA index: 374
type NodeChroma struct {
	T1        float32
	T2        float32
	T3        float32
	Fsize     float32
	Fstrength float32
	Falpha    float32
	Key       [4]float32
	Algorithm int16
	Channel   int16
}

// SDNA index: 375
type NodeTwoXYs struct {
	X1     int16
	X2     int16
	Y1     int16
	Y2     int16
	Fac_x1 float32
	Fac_x2 float32
	Fac_y1 float32
	Fac_y2 float32
}

// SDNA index: 376
type NodeTwoFloats struct {
	X float32
	Y float32
}

// SDNA index: 377
type NodeGeometry struct {
	Uvname  [64]int8
	Colname [64]int8
}

// SDNA index: 378
type NodeVertexCol struct {
	Name [64]int8
}

// SDNA index: 379
type NodeDefocus struct {
	Bktype   int8
	Pad_c1   int8
	Preview  int8
	Gamco    int8
	Samples  int16
	No_zbuf  int16
	Fstop    float32
	Maxblur  float32
	Bthresh  float32
	Scale    float32
	Rotation float32
	Pad_f1   float32
}

// SDNA index: 380
type NodeScriptDict struct {
	Dict Pointer // *struct{}
	Node Pointer // *struct{}
}

// SDNA index: 381
type NodeGlare struct {
	Quality   int8
	Type      int8
	Iter      int8
	Angle     int8
	Pad_c1    int8
	Size      int8
	Pad       [2]int8
	Colmod    float32
	Mix       float32
	Threshold float32
	Fade      float32
	Angle_ofs float32
	Pad_f1    float32
}

// SDNA index: 382
type NodeTonemap struct {
	Key    float32
	Offset float32
	Gamma  float32
	F      float32
	M      float32
	A      float32
	C      float32
	Type   int32
}

// SDNA index: 383
type NodeLensDist struct {
	Jit  int16
	Proj int16
	Fit  int16
	Pad  int16
}

// SDNA index: 384
type NodeColorBalance struct {
	Slope     [3]float32
	Offset    [3]float32
	Power     [3]float32
	Lift      [3]float32
	Gamma     [3]float32
	Gain      [3]float32
	Lift_lgg  [3]float32
	Gamma_inv [3]float32
}

// SDNA index: 385
type NodeColorspill struct {
	Limchan  int16
	Unspill  int16
	Limscale float32
	Uspillr  float32
	Uspillg  float32
	Uspillb  float32
}

// SDNA index: 386
type NodeDilateErode struct {
	Falloff int8
	Pad     [7]int8
}

// SDNA index: 387
type NodeMask struct {
	Size_x int32
	Size_y int32
}

// SDNA index: 388
type NodeTexBase struct {
	Tex_mapping   TexMapping
	Color_mapping ColorMapping
}

// SDNA index: 389
type NodeTexSky struct {
	Base          NodeTexBase
	Sun_direction [3]float32
	Turbidity     float32
}

// SDNA index: 390
type NodeTexImage struct {
	Base             NodeTexBase
	Iuser            ImageUser
	Color_space      int32
	Projection       int32
	Projection_blend float32
	Pad              int32
}

// SDNA index: 391
type NodeTexChecker struct {
	Base NodeTexBase
}

// SDNA index: 392
type NodeTexBrick struct {
	Base        NodeTexBase
	Offset_freq int32
	Squash_freq int32
	Offset      float32
	Squash      float32
}

// SDNA index: 393
type NodeTexEnvironment struct {
	Base        NodeTexBase
	Iuser       ImageUser
	Color_space int32
	Projection  int32
}

// SDNA index: 394
type NodeTexGradient struct {
	Base          NodeTexBase
	Gradient_type int32
	Pad           int32
}

// SDNA index: 395
type NodeTexNoise struct {
	Base NodeTexBase
}

// SDNA index: 396
type NodeTexVoronoi struct {
	Base     NodeTexBase
	Coloring int32
	Pad      int32
}

// SDNA index: 397
type NodeTexMusgrave struct {
	Base          NodeTexBase
	Musgrave_type int32
	Pad           int32
}

// SDNA index: 398
type NodeTexWave struct {
	Base      NodeTexBase
	Wave_type int32
	Pad       int32
}

// SDNA index: 399
type NodeTexMagic struct {
	Base  NodeTexBase
	Depth int32
	Pad   int32
}

// SDNA index: 400
type NodeShaderAttribute struct {
	Name [64]int8
}

// SDNA index: 401
type TexNodeOutput struct {
	Name [64]int8
}

// SDNA index: 402
type NodeKeyingScreenData struct {
	Tracking_object [64]int8
}

// SDNA index: 403
type NodeKeyingData struct {
	Screen_balance        float32
	Despill_factor        float32
	Despill_balance       float32
	Edge_kernel_radius    int32
	Edge_kernel_tolerance float32
	Clip_black            float32
	Clip_white            float32
	Dilate_distance       int32
	Feather_distance      int32
	Feather_falloff       int32
	Blur_pre              int32
	Blur_post             int32
}

// SDNA index: 404
type NodeTrackPosData struct {
	Tracking_object [64]int8
	Track_name      [64]int8
}

// SDNA index: 405
type NodeTranslateData struct {
	Wrap_axis int8
	Relative  int8
	Pad       [6]int8
}

// SDNA index: 406
type NodeShaderScript struct {
	Mode          int32
	Flag          int32
	Filepath      [1024]int8
	Bytecode_hash [64]int8
	Bytecode      Pointer // *int8
}

// SDNA index: 407
type NodeShaderTangent struct {
	Direction_type int32
	Axis           int32
	Uv_map         [64]int8
}

// SDNA index: 408
type NodeShaderNormalMap struct {
	Space  int32
	Uv_map [64]int8
}

// SDNA index: 409
type CurveMapPoint struct {
	X      float32
	Y      float32
	Flag   int16
	Shorty int16
}

// SDNA index: 410
type CurveMap struct {
	Totpoint    int16
	Flag        int16
	Range       float32
	Mintable    float32
	Maxtable    float32
	Ext_in      [2]float32
	Ext_out     [2]float32
	Curve       Pointer // *CurveMapPoint
	Table       Pointer // *CurveMapPoint
	Premultable Pointer // *CurveMapPoint
}

// SDNA index: 411
type CurveMapping struct {
	Flag              int32
	Cur               int32
	Preset            int32
	Changed_timestamp int32
	Curr              Rctf
	Clipr             Rctf
	Cm                [4]CurveMap
	Black             [3]float32
	White             [3]float32
	Bwmul             [3]float32
	Sample            [3]float32
}

// SDNA index: 412
type Histogram struct {
	Channels     int32
	X_resolution int32
	Data_luma    [256]float32
	Data_r       [256]float32
	Data_g       [256]float32
	Data_b       [256]float32
	Data_a       [256]float32
	Xmax         float32
	Ymax         float32
	Mode         int16
	Flag         int16
	Height       int32
	Co           [2][2]float32
}

// SDNA index: 413
type Scopes struct {
	Ok              int32
	Sample_full     int32
	Sample_lines    int32
	Accuracy        float32
	Wavefrm_mode    int32
	Wavefrm_alpha   float32
	Wavefrm_yfac    float32
	Wavefrm_height  int32
	Vecscope_alpha  float32
	Vecscope_height int32
	Minmax          [3][2]float32
	Hist            Histogram
	Waveform_1      Pointer // *float32
	Waveform_2      Pointer // *float32
	Waveform_3      Pointer // *float32
	Vecscope        Pointer // *float32
	Waveform_tot    int32
	Pad             int32
}

// SDNA index: 414
type ColorManagedViewSettings struct {
	Flag           int32
	Pad            int32
	View_transform [64]int8
	Exposure       float32
	Gamma          float32
	Curve_mapping  Pointer // *CurveMapping
	Pad2           Pointer // *struct{}
}

// SDNA index: 415
type ColorManagedDisplaySettings struct {
	Display_device [64]int8
}

// SDNA index: 416
type ColorManagedColorspaceSettings struct {
	Name [64]int8
}

// SDNA index: 417
type BrushClone struct {
	Image  Pointer // *Image
	Offset [2]float32
	Alpha  float32
	Pad    float32
}

// SDNA index: 418
type Brush struct {
	Id                     ID
	Clone                  BrushClone
	Curve                  Pointer // *CurveMapping
	Mtex                   MTex
	Mask_mtex              MTex
	Toggle_brush           Pointer // *Brush
	Icon_imbuf             Pointer // *ImBuf
	Preview                Pointer // *PreviewImage
	Icon_filepath          [1024]int8
	Normal_weight          float32
	Blend                  int16
	Ob_mode                int16
	Weight                 float32
	Size                   int32
	Flag                   int32
	Jitter                 float32
	Jitter_absolute        int32
	Overlay_flags          int32
	Spacing                int32
	Smooth_stroke_radius   int32
	Smooth_stroke_factor   float32
	Rate                   float32
	Rgb                    [3]float32
	Alpha                  float32
	Sculpt_plane           int32
	Plane_offset           float32
	Sculpt_tool            int8
	Vertexpaint_tool       int8
	Imagepaint_tool        int8
	Mask_tool              int8
	Autosmooth_factor      float32
	Crease_pinch_factor    float32
	Plane_trim             float32
	Height                 float32
	Texture_sample_bias    float32
	Texture_overlay_alpha  int32
	Mask_overlay_alpha     int32
	Cursor_overlay_alpha   int32
	Unprojected_radius     float32
	Add_col                [3]float32
	Sub_col                [3]float32
	Stencil_pos            [2]float32
	Stencil_dimension      [2]float32
	Mask_stencil_pos       [2]float32
	Mask_stencil_dimension [2]float32
}

// SDNA index: 419
type CustomDataLayer struct {
	Type         int32
	Offset       int32
	Flag         int32
	Active       int32
	Active_rnd   int32
	Active_clone int32
	Active_mask  int32
	Uid          int32
	Name         [64]int8
	Data         Pointer // *struct{}
}

// SDNA index: 420
type CustomDataExternal struct {
	Filename [1024]int8
}

// SDNA index: 421
type CustomData struct {
	Layers   Pointer // *CustomDataLayer
	Typemap  [39]int32
	Totlayer int32
	Maxlayer int32
	Totsize  int32
	Pool     Pointer // *struct{}
	External Pointer // *CustomDataExternal
}

// SDNA index: 422
type HairKey struct {
	Co       [3]float32
	Time     float32
	Weight   float32
	Editflag int16
	Pad      int16
}

// SDNA index: 423
type ParticleKey struct {
	Co   [3]float32
	Vel  [3]float32
	Rot  [4]float32
	Ave  [3]float32
	Time float32
}

// SDNA index: 424
type BoidParticle struct {
	Ground  Pointer // *Object
	Data    BoidData
	Gravity [3]float32
	Wander  [3]float32
	Rt      float32
}

// SDNA index: 425
type ParticleSpring struct {
	Rest_length    float32
	Particle_index [2]int32
	Delete_flag    int32
}

// SDNA index: 426
type ChildParticle struct {
	Num     int32
	Parent  int32
	Pa      [4]int32
	W       [4]float32
	Fuv     [4]float32
	Foffset float32
	Rt      float32
}

// SDNA index: 427
type ParticleTarget struct {
	Next     Pointer // *ParticleTarget
	Prev     Pointer // *ParticleTarget
	Ob       Pointer // *Object
	Psys     int32
	Flag     int16
	Mode     int16
	Time     float32
	Duration float32
}

// SDNA index: 428
type ParticleDupliWeight struct {
	Next  Pointer // *ParticleDupliWeight
	Prev  Pointer // *ParticleDupliWeight
	Ob    Pointer // *Object
	Count int16
	Flag  int16
	Index int16
	Rt    int16
}

// SDNA index: 429
type ParticleData struct {
	State       ParticleKey
	Prev_state  ParticleKey
	Hair        Pointer // *HairKey
	Keys        Pointer // *ParticleKey
	Boid        Pointer // *BoidParticle
	Totkey      int32
	Time        float32
	Lifetime    float32
	Dietime     float32
	Num         int32
	Num_dmcache int32
	Fuv         [4]float32
	Foffset     float32
	Size        float32
	Sphdensity  float32
	Pad         int32
	Hair_index  int32
	Flag        int16
	Alive       int16
}

// SDNA index: 430
type SPHFluidSettings struct {
	Radius              float32
	Spring_k            float32
	Rest_length         float32
	Plasticity_constant float32
	Yield_ratio         float32
	Plasticity_balance  float32
	Yield_balance       float32
	Viscosity_omega     float32
	Viscosity_beta      float32
	Stiffness_k         float32
	Stiffness_knear     float32
	Rest_density        float32
	Buoyancy            float32
	Flag                int32
	Spring_frames       int32
	Solver              int16
	Pad                 [3]int16
}

// SDNA index: 431
type ParticleSettings struct {
	Id                  ID
	Adt                 Pointer // *AnimData
	Boids               Pointer // *BoidSettings
	Fluid               Pointer // *SPHFluidSettings
	Effector_weights    Pointer // *EffectorWeights
	Flag                int32
	Rt                  int32
	Type                int16
	From                int16
	Distr               int16
	Texact              int16
	Phystype            int16
	Rotmode             int16
	Avemode             int16
	Reactevent          int16
	Draw                int32
	Pad1                int32
	Draw_as             int16
	Draw_size           int16
	Childtype           int16
	Pad2                int16
	Ren_as              int16
	Subframes           int16
	Draw_col            int16
	Draw_step           int16
	Ren_step            int16
	Hair_step           int16
	Keys_step           int16
	Adapt_angle         int16
	Adapt_pix           int16
	Disp                int16
	Omat                int16
	Interpolation       int16
	Integrator          int16
	Rotfrom             int16
	Kink                int16
	Kink_axis           int16
	Bb_align            int16
	Bb_uv_split         int16
	Bb_anim             int16
	Bb_split_offset     int16
	Bb_tilt             float32
	Bb_rand_tilt        float32
	Bb_offset           [2]float32
	Bb_size             [2]float32
	Bb_vel_head         float32
	Bb_vel_tail         float32
	Color_vec_max       float32
	Simplify_flag       int16
	Simplify_refsize    int16
	Simplify_rate       float32
	Simplify_transition float32
	Simplify_viewport   float32
	Sta                 float32
	End                 float32
	Lifetime            float32
	Randlife            float32
	Timetweak           float32
	Courant_target      float32
	Jitfac              float32
	Eff_hair            float32
	Grid_rand           float32
	Ps_offset           [1]float32
	Totpart             int32
	Userjit             int32
	Grid_res            int32
	Effector_amount     int32
	Time_flag           int16
	Time_pad            [3]int16
	Normfac             float32
	Obfac               float32
	Randfac             float32
	Partfac             float32
	Tanfac              float32
	Tanphase            float32
	Reactfac            float32
	Ob_vel              [3]float32
	Avefac              float32
	Phasefac            float32
	Randrotfac          float32
	Randphasefac        float32
	Mass                float32
	Size                float32
	Randsize            float32
	Acc                 [3]float32
	Dragfac             float32
	Brownfac            float32
	Dampfac             float32
	Randlength          float32
	Child_nbr           int32
	Ren_child_nbr       int32
	Parents             float32
	Childsize           float32
	Childrandsize       float32
	Childrad            float32
	Childflat           float32
	Clumpfac            float32
	Clumppow            float32
	Kink_amp            float32
	Kink_freq           float32
	Kink_shape          float32
	Kink_flat           float32
	Kink_amp_clump      float32
	Rough1              float32
	Rough1_size         float32
	Rough2              float32
	Rough2_size         float32
	Rough2_thres        float32
	Rough_end           float32
	Rough_end_shape     float32
	Clength             float32
	Clength_thres       float32
	Parting_fac         float32
	Parting_min         float32
	Parting_max         float32
	Branch_thres        float32
	Draw_line           [2]float32
	Path_start          float32
	Path_end            float32
	Trail_count         int32
	Keyed_loops         int32
	Mtex                [18]Pointer // [18]*MTex
	Dup_group           Pointer     // *Group
	Dupliweights        ListBase
	Eff_group           Pointer // *Group
	Dup_ob              Pointer // *Object
	Bb_ob               Pointer // *Object
	Ipo                 Pointer // *Ipo
	Pd                  Pointer // *PartDeflect
	Pd2                 Pointer // *PartDeflect
}

// SDNA index: 432
type ParticleSystem struct {
	Next               Pointer // *ParticleSystem
	Prev               Pointer // *ParticleSystem
	Part               Pointer // *ParticleSettings
	Particles          Pointer // *ParticleData
	Child              Pointer // *ChildParticle
	Edit               Pointer // *PTCacheEdit
	Free_edit          func()
	Pathcache          Pointer // **ParticleCacheKey
	Childcache         Pointer // **ParticleCacheKey
	Pathcachebufs      ListBase
	Childcachebufs     ListBase
	Clmd               Pointer // *ClothModifierData
	Hair_in_dm         Pointer // *DerivedMesh
	Hair_out_dm        Pointer // *DerivedMesh
	Target_ob          Pointer // *Object
	Lattice            Pointer // *Object
	Parent             Pointer // *Object
	Targets            ListBase
	Name               [64]int8
	Imat               [4][4]float32
	Cfra               float32
	Tree_frame         float32
	Bvhtree_frame      float32
	Seed               int32
	Child_seed         int32
	Flag               int32
	Totpart            int32
	Totunexist         int32
	Totchild           int32
	Totcached          int32
	Totchildcache      int32
	Recalc             int16
	Target_psys        int16
	Totkeyed           int16
	Bakespace          int16
	Bb_uvname          [3][64]int8
	Vgroup             [12]int16
	Vg_neg             int16
	Rt3                int16
	Renderdata         Pointer // *struct{}
	Pointcache         Pointer // *PointCache
	Ptcaches           ListBase
	Effectors          Pointer // *ListBase
	Fluid_springs      Pointer // *ParticleSpring
	Tot_fluidsprings   int32
	Alloc_fluidsprings int32
	Tree               Pointer // *KDTree
	Bvhtree            Pointer // *BVHTree
	Pdd                Pointer // *ParticleDrawData
	Frand              Pointer // *float32
	Dt_frac            float32
	_pad               float32
}

// SDNA index: 433
type ClothSimSettings struct {
	Cache             Pointer // *LinkNode
	Mingoal           float32
	Cdis              float32
	Cvi               float32
	Gravity           [3]float32
	Dt                float32
	Mass              float32
	Structural        float32
	Shear             float32
	Bending           float32
	Max_bend          float32
	Max_struct        float32
	Max_shear         float32
	Avg_spring_len    float32
	Timescale         float32
	Maxgoal           float32
	Eff_force_scale   float32
	Eff_wind_scale    float32
	Sim_time_old      float32
	Defgoal           float32
	Goalspring        float32
	Goalfrict         float32
	Velocity_smooth   float32
	Collider_friction float32
	Vel_damping       float32
	StepsPerFrame     int32
	Flags             int32
	Preroll           int32
	Maxspringlen      int32
	Solver_type       int16
	Vgroup_bend       int16
	Vgroup_mass       int16
	Vgroup_struct     int16
	Shapekey_rest     int16
	Presets           int16
	Reset             int16
	Pad               int16
	Effector_weights  Pointer // *EffectorWeights
}

// SDNA index: 434
type ClothCollSettings struct {
	Collision_list  Pointer // *LinkNode
	Epsilon         float32
	Self_friction   float32
	Friction        float32
	Selfepsilon     float32
	Repel_force     float32
	Distance_repel  float32
	Flags           int32
	Self_loop_count int16
	Loop_count      int16
	Group           Pointer // *Group
	Vgroup_selfcol  int16
	Pad             int16
	Pad2            int32
}

// SDNA index: 435
type BGPDspoint struct {
	X        float32
	Y        float32
	Z        float32
	Pressure float32
	Time     float32
}

// SDNA index: 436
type BGPDstroke struct {
	Next      Pointer // *BGPDstroke
	Prev      Pointer // *BGPDstroke
	Points    Pointer // *BGPDspoint
	Pad       Pointer // *struct{}
	Totpoints int32
	Thickness int16
	Flag      int16
	Inittime  float64
}

// SDNA index: 437
type BGPDframe struct {
	Next     Pointer // *BGPDframe
	Prev     Pointer // *BGPDframe
	Strokes  ListBase
	Framenum int32
	Flag     int32
}

// SDNA index: 438
type BGPDlayer struct {
	Next      Pointer // *BGPDlayer
	Prev      Pointer // *BGPDlayer
	Frames    ListBase
	Actframe  Pointer // *BGPDframe
	Flag      int32
	Thickness int16
	Gstep     int16
	Color     [4]float32
	Info      [128]int8
}

// SDNA index: 439
type BGPdata struct {
	Id            ID
	Layers        ListBase
	Flag          int32
	Sbuffer_size  int16
	Sbuffer_sflag int16
	Sbuffer       Pointer // *struct{}
}

// SDNA index: 440
type ReportList struct {
	List        ListBase
	Printlevel  int32
	Storelevel  int32
	Flag        int32
	Pad         int32
	Reporttimer Pointer // *WmTimer
}

// SDNA index: 441
type WmWindowManager struct {
	Id            ID
	Windrawable   Pointer // *WmWindow
	Winactive     Pointer // *WmWindow
	Windows       ListBase
	Initialized   int32
	File_saved    int16
	Op_undo_depth int16
	Operators     ListBase
	Queue         ListBase
	Reports       ReportList
	Jobs          ListBase
	Paintcursors  ListBase
	Drags         ListBase
	Keyconfigs    ListBase
	Defaultconf   Pointer // *WmKeyConfig
	Addonconf     Pointer // *WmKeyConfig
	Userconf      Pointer // *WmKeyConfig
	Timers        ListBase
	Autosavetimer Pointer // *WmTimer
}

// SDNA index: 442
type WmWindow struct {
	Next          Pointer // *WmWindow
	Prev          Pointer // *WmWindow
	Ghostwin      Pointer // *struct{}
	Winid         int32
	Grabcursor    int16
	Pad           int16
	Screen        Pointer // *BScreen
	Newscreen     Pointer // *BScreen
	Screenname    [64]int8
	Posx          int16
	Posy          int16
	Sizex         int16
	Sizey         int16
	Windowstate   int16
	Monitor       int16
	Active        int16
	Cursor        int16
	Lastcursor    int16
	Modalcursor   int16
	Addmousemove  int16
	Pad2          int16
	Eventstate    Pointer // *WmEvent
	Curswin       Pointer // *WmSubWindow
	Tweak         Pointer // *WmGesture
	Drawmethod    int32
	Drawfail      int32
	Drawdata      Pointer // *struct{}
	Queue         ListBase
	Handlers      ListBase
	Modalhandlers ListBase
	Subwindows    ListBase
	Gesture       ListBase
}

// SDNA index: 443
type WmKeyMapItem struct {
	Next          Pointer // *WmKeyMapItem
	Prev          Pointer // *WmKeyMapItem
	Idname        [64]int8
	Properties    Pointer // *IDProperty
	Propvalue_str [64]int8
	Propvalue     int16
	Type          int16
	Val           int16
	Shift         int16
	Ctrl          int16
	Alt           int16
	Oskey         int16
	Keymodifier   int16
	Flag          int16
	Maptype       int16
	Id            int16
	Pad           int16
	Ptr           Pointer // *PointerRNA
}

// SDNA index: 444
type WmKeyMapDiffItem struct {
	Next        Pointer // *WmKeyMapDiffItem
	Prev        Pointer // *WmKeyMapDiffItem
	Remove_item Pointer // *WmKeyMapItem
	Add_item    Pointer // *WmKeyMapItem
}

// SDNA index: 445
type WmKeyMap struct {
	Next        Pointer // *WmKeyMap
	Prev        Pointer // *WmKeyMap
	Items       ListBase
	Diff_items  ListBase
	Idname      [64]int8
	Spaceid     int16
	Regionid    int16
	Flag        int16
	Kmi_id      int16
	Poll        func() int32
	Modal_items Pointer // *struct{}
}

// SDNA index: 446
type WmKeyConfig struct {
	Next      Pointer // *WmKeyConfig
	Prev      Pointer // *WmKeyConfig
	Idname    [64]int8
	Basename  [64]int8
	Keymaps   ListBase
	Actkeymap int32
	Flag      int32
}

// SDNA index: 447
type WmOperator struct {
	Next        Pointer // *WmOperator
	Prev        Pointer // *WmOperator
	Idname      [64]int8
	Properties  Pointer // *IDProperty
	Type        Pointer // *WmOperatorType
	Customdata  Pointer // *struct{}
	Py_instance Pointer // *struct{}
	Ptr         Pointer // *PointerRNA
	Reports     Pointer // *ReportList
	Macro       ListBase
	Opm         Pointer // *WmOperator
	Layout      Pointer // *UiLayout
	Flag        int16
	Pad         [3]int16
}

// SDNA index: 448
type FModifier struct {
	Next      Pointer // *FModifier
	Prev      Pointer // *FModifier
	Data      Pointer // *struct{}
	Edata     Pointer // *struct{}
	Name      [64]int8
	Type      int16
	Flag      int16
	Influence float32
	Sfra      float32
	Efra      float32
	Blendin   float32
	Blendout  float32
}

// SDNA index: 449
type FMod_Generator struct {
	Coefficients Pointer // *float32
	Arraysize    int32
	Poly_order   int32
	Mode         int32
	Flag         int32
}

// SDNA index: 450
type FMod_FunctionGenerator struct {
	Amplitude        float32
	Phase_multiplier float32
	Phase_offset     float32
	Value_offset     float32
	Type             int32
	Flag             int32
}

// SDNA index: 451
type FCM_EnvelopeData struct {
	Min  float32
	Max  float32
	Time float32
	F1   int16
	F2   int16
}

// SDNA index: 452
type FMod_Envelope struct {
	Data    Pointer // *FCM_EnvelopeData
	Totvert int32
	Midval  float32
	Min     float32
	Max     float32
}

// SDNA index: 453
type FMod_Cycles struct {
	Before_mode   int16
	After_mode    int16
	Before_cycles int16
	After_cycles  int16
}

// SDNA index: 454
type FMod_Python struct {
	Script Pointer // *Text
	Prop   Pointer // *IDProperty
}

// SDNA index: 455
type FMod_Limits struct {
	Rect Rctf
	Flag int32
	Pad  int32
}

// SDNA index: 456
type FMod_Noise struct {
	Size         float32
	Strength     float32
	Phase        float32
	Pad          float32
	Depth        int16
	Modification int16
}

// SDNA index: 457
type FMod_Stepped struct {
	Step_size   float32
	Offset      float32
	Start_frame float32
	End_frame   float32
	Flag        int32
}

// SDNA index: 458
type DriverTarget struct {
	Id         Pointer // *ID
	Rna_path   Pointer // *int8
	Pchan_name [32]int8
	TransChan  int16
	Flag       int16
	Idtype     int32
}

// SDNA index: 459
type DriverVar struct {
	Next        Pointer // *DriverVar
	Prev        Pointer // *DriverVar
	Name        [64]int8
	Targets     [8]DriverTarget
	Num_targets int16
	Type        int16
	Curval      float32
}

// SDNA index: 460
type ChannelDriver struct {
	Variables  ListBase
	Expression [256]int8
	Expr_comp  Pointer // *struct{}
	Curval     float32
	Influence  float32
	Type       int32
	Flag       int32
}

// SDNA index: 461
type FPoint struct {
	Vec  [2]float32
	Flag int32
	Pad  int32
}

// SDNA index: 462
type FCurve struct {
	Next        Pointer // *FCurve
	Prev        Pointer // *FCurve
	Grp         Pointer // *BActionGroup
	Driver      Pointer // *ChannelDriver
	Modifiers   ListBase
	Bezt        Pointer // *BezTriple
	Fpt         Pointer // *FPoint
	Totvert     int32
	Curval      float32
	Flag        int16
	Extend      int16
	Array_index int32
	Rna_path    Pointer // *int8
	Color_mode  int32
	Color       [3]float32
}

// SDNA index: 463
type AnimMapPair struct {
	From [128]int8
	To   [128]int8
}

// SDNA index: 464
type AnimMapper struct {
	Next     Pointer // *AnimMapper
	Prev     Pointer // *AnimMapper
	Target   Pointer // *BAction
	Mappings ListBase
}

// SDNA index: 465
type NlaStrip struct {
	Next           Pointer // *NlaStrip
	Prev           Pointer // *NlaStrip
	Strips         ListBase
	Act            Pointer // *BAction
	Remap          Pointer // *AnimMapper
	Fcurves        ListBase
	Modifiers      ListBase
	Name           [64]int8
	Influence      float32
	Strip_time     float32
	Start          float32
	End            float32
	Actstart       float32
	Actend         float32
	Repeat         float32
	Scale          float32
	Blendin        float32
	Blendout       float32
	Blendmode      int16
	Extendmode     int16
	Pad1           int16
	Type           int16
	Speaker_handle Pointer // *struct{}
	Flag           int32
	Pad2           int32
}

// SDNA index: 466
type NlaTrack struct {
	Next   Pointer // *NlaTrack
	Prev   Pointer // *NlaTrack
	Strips ListBase
	Flag   int32
	Index  int32
	Name   [64]int8
}

// SDNA index: 467
type KS_Path struct {
	Next        Pointer // *KS_Path
	Prev        Pointer // *KS_Path
	Id          Pointer // *ID
	Group       [64]int8
	Idtype      int32
	Groupmode   int16
	Pad         int16
	Rna_path    Pointer // *int8
	Array_index int32
	Flag        int16
	Keyingflag  int16
}

// SDNA index: 468
type KeyingSet struct {
	Next        Pointer // *KeyingSet
	Prev        Pointer // *KeyingSet
	Paths       ListBase
	Idname      [64]int8
	Name        [64]int8
	Description [240]int8
	Typeinfo    [64]int8
	Flag        int16
	Keyingflag  int16
	Active_path int32
}

// SDNA index: 469
type AnimOverride struct {
	Next        Pointer // *AnimOverride
	Prev        Pointer // *AnimOverride
	Rna_path    Pointer // *int8
	Array_index int32
	Value       float32
}

// SDNA index: 470
type AnimData struct {
	Action         Pointer // *BAction
	Tmpact         Pointer // *BAction
	Remap          Pointer // *AnimMapper
	Nla_tracks     ListBase
	Actstrip       Pointer // *NlaStrip
	Drivers        ListBase
	Overrides      ListBase
	Flag           int32
	Recalc         int32
	Act_blendmode  int16
	Act_extendmode int16
	Act_influence  float32
}

// SDNA index: 471
type IdAdtTemplate struct {
	Id  ID
	Adt Pointer // *AnimData
}

// SDNA index: 472
type BoidRule struct {
	Next Pointer // *BoidRule
	Prev Pointer // *BoidRule
	Type int32
	Flag int32
	Name [32]int8
}

// SDNA index: 473
type BoidRuleGoalAvoid struct {
	Rule        BoidRule
	Ob          Pointer // *Object
	Options     int32
	Fear_factor float32
	Signal_id   int32
	Channels    int32
}

// SDNA index: 474
type BoidRuleAvoidCollision struct {
	Rule       BoidRule
	Options    int32
	Look_ahead float32
}

// SDNA index: 475
type BoidRuleFollowLeader struct {
	Rule       BoidRule
	Ob         Pointer // *Object
	Loc        [3]float32
	Oloc       [3]float32
	Cfra       float32
	Distance   float32
	Options    int32
	Queue_size int32
}

// SDNA index: 476
type BoidRuleAverageSpeed struct {
	Rule   BoidRule
	Wander float32
	Level  float32
	Speed  float32
	Rt     float32
}

// SDNA index: 477
type BoidRuleFight struct {
	Rule          BoidRule
	Distance      float32
	Flee_distance float32
}

// SDNA index: 478
type BoidData struct {
	Health   float32
	Acc      [3]float32
	State_id int16
	Mode     int16
}

// SDNA index: 479
type BoidState struct {
	Next           Pointer // *BoidState
	Prev           Pointer // *BoidState
	Rules          ListBase
	Conditions     ListBase
	Actions        ListBase
	Name           [32]int8
	Id             int32
	Flag           int32
	Ruleset_type   int32
	Rule_fuzziness float32
	Signal_id      int32
	Channels       int32
	Volume         float32
	Falloff        float32
}

// SDNA index: 480
type BoidSettings struct {
	Options             int32
	Last_state_id       int32
	Landing_smoothness  float32
	Height              float32
	Banking             float32
	Pitch               float32
	Health              float32
	Aggression          float32
	Strength            float32
	Accuracy            float32
	Range               float32
	Air_min_speed       float32
	Air_max_speed       float32
	Air_max_acc         float32
	Air_max_ave         float32
	Air_personal_space  float32
	Land_jump_speed     float32
	Land_max_speed      float32
	Land_max_acc        float32
	Land_max_ave        float32
	Land_personal_space float32
	Land_stick_force    float32
	States              ListBase
}

// SDNA index: 481
type SmokeDomainSettings struct {
	Smd               Pointer // *SmokeModifierData
	Fluid             Pointer // *FLUID_3D
	Fluid_mutex       Pointer // *struct{}
	Fluid_group       Pointer // *Group
	Eff_group         Pointer // *Group
	Coll_group        Pointer // *Group
	Wt                Pointer // *WTURBULENCE
	Tex               Pointer // *GPUTexture
	Tex_wt            Pointer // *GPUTexture
	Tex_shadow        Pointer // *GPUTexture
	Tex_flame         Pointer // *GPUTexture
	Shadow            Pointer // *float32
	P0                [3]float32
	P1                [3]float32
	Dp0               [3]float32
	Cell_size         [3]float32
	Global_size       [3]float32
	Prev_loc          [3]float32
	Shift             [3]int32
	Shift_f           [3]float32
	Obj_shift_f       [3]float32
	Imat              [4][4]float32
	Obmat             [4][4]float32
	Base_res          [3]int32
	Res_min           [3]int32
	Res_max           [3]int32
	Res               [3]int32
	Total_cells       int32
	Dx                float32
	Scale             float32
	Adapt_margin      int32
	Adapt_res         int32
	Adapt_threshold   float32
	Alpha             float32
	Beta              float32
	Amplify           int32
	Maxres            int32
	Flags             int32
	Viewsettings      int32
	Noise             int16
	Diss_percent      int16
	Diss_speed        int32
	Strength          float32
	Res_wt            [3]int32
	Dx_wt             float32
	Cache_comp        int32
	Cache_high_comp   int32
	Point_cache       [2]Pointer // [2]*PointCache
	Ptcaches          [2]ListBase
	Effector_weights  Pointer // *EffectorWeights
	Border_collisions int32
	Time_scale        float32
	Vorticity         float32
	Active_fields     int32
	Active_color      [3]float32
	Pad               int32
	Burning_rate      float32
	Flame_smoke       float32
	Flame_vorticity   float32
	Flame_ignition    float32
	Flame_max_temp    float32
	Flame_smoke_color [3]float32
}

// SDNA index: 482
type SmokeFlowSettings struct {
	Smd              Pointer // *SmokeModifierData
	Dm               Pointer // *DerivedMesh
	Psys             Pointer // *ParticleSystem
	Noise_texture    Pointer // *Tex
	Verts_old        Pointer // *float32
	Numverts         int32
	Vel_multi        float32
	Vel_normal       float32
	Vel_random       float32
	Density          float32
	Color            [3]float32
	Fuel_amount      float32
	Temp             float32
	Volume_density   float32
	Surface_distance float32
	Texture_size     float32
	Texture_offset   float32
	Pad              int32
	Uvlayer_name     [64]int8
	Vgroup_density   int16
	Type             int16
	Source           int16
	Texture_type     int16
	Flags            int32
}

// SDNA index: 483
type SmokeCollSettings struct {
	Smd       Pointer // *SmokeModifierData
	Dm        Pointer // *DerivedMesh
	Verts_old Pointer // *float32
	Numverts  int32
	Type      int16
	Pad       int16
}

// SDNA index: 484
type Speaker struct {
	Id                 ID
	Adt                Pointer // *AnimData
	Sound              Pointer // *BSound
	Volume_max         float32
	Volume_min         float32
	Distance_max       float32
	Distance_reference float32
	Attenuation        float32
	Cone_angle_outer   float32
	Cone_angle_inner   float32
	Cone_volume_outer  float32
	Volume             float32
	Pitch              float32
	Flag               int16
	Pad1               [3]int16
}

// SDNA index: 485
type MovieClipUser struct {
	Framenr     int32
	Render_size int16
	Render_flag int16
}

// SDNA index: 486
type MovieClipProxy struct {
	Dir             [768]int8
	Tc              int16
	Quality         int16
	Build_size_flag int16
	Build_tc_flag   int16
}

// SDNA index: 487
type MovieClip struct {
	Id                  ID
	Adt                 Pointer // *AnimData
	Name                [1024]int8
	Source              int32
	Lastframe           int32
	Lastsize            [2]int32
	Aspx                float32
	Aspy                float32
	Anim                Pointer // *Anim
	Cache               Pointer // *MovieClipCache
	Gpd                 Pointer // *BGPdata
	Tracking            MovieTracking
	Tracking_context    Pointer // *struct{}
	Proxy               MovieClipProxy
	Flag                int32
	Len                 int32
	Start_frame         int32
	Frame_offset        int32
	Colorspace_settings ColorManagedColorspaceSettings
}

// SDNA index: 488
type MovieClipScopes struct {
	Ok                   int16
	Use_track_mask       int16
	Track_preview_height int32
	Frame_width          int32
	Frame_height         int32
	Undist_marker        MovieTrackingMarker
	Track_search         Pointer // *ImBuf
	Track_preview        Pointer // *ImBuf
	Track_pos            [2]float32
	Track_disabled       int16
	Track_locked         int16
	Framenr              int32
	Track                Pointer // *MovieTrackingTrack
	Marker               Pointer // *MovieTrackingMarker
	Slide_scale          [2]float32
}

// SDNA index: 489
type MovieReconstructedCamera struct {
	Framenr int32
	Error   float32
	Mat     [4][4]float32
}

// SDNA index: 490
type MovieTrackingCamera struct {
	Intrinsics   Pointer // *struct{}
	Sensor_width float32
	Pixel_aspect float32
	Pad          float32
	Focal        float32
	Units        int16
	Pad1         int16
	Principal    [2]float32
	K1           float32
	K2           float32
	K3           float32
}

// SDNA index: 491
type MovieTrackingMarker struct {
	Pos             [2]float32
	Pattern_corners [4][2]float32
	Search_min      [2]float32
	Search_max      [2]float32
	Framenr         int32
	Flag            int32
}

// SDNA index: 492
type MovieTrackingTrack struct {
	Next                Pointer // *MovieTrackingTrack
	Prev                Pointer // *MovieTrackingTrack
	Name                [64]int8
	Pat_min             [2]float32
	Pat_max             [2]float32
	Search_min          [2]float32
	Search_max          [2]float32
	Offset              [2]float32
	Markersnr           int32
	Last_marker         int32
	Markers             Pointer // *MovieTrackingMarker
	Bundle_pos          [3]float32
	Error               float32
	Flag                int32
	Pat_flag            int32
	Search_flag         int32
	Color               [3]float32
	Frames_limit        int16
	Margin              int16
	Pattern_match       int16
	Motion_model        int16
	Algorithm_flag      int32
	Minimum_correlation float32
	Gpd                 Pointer // *BGPdata
}

// SDNA index: 493
type MovieTrackingSettings struct {
	Flag                             int32
	Default_motion_model             int16
	Default_algorithm_flag           int16
	Default_minimum_correlation      float32
	Default_pattern_size             int16
	Default_search_size              int16
	Default_frames_limit             int16
	Default_margin                   int16
	Default_pattern_match            int16
	Default_flag                     int16
	Motion_flag                      int16
	Speed                            int16
	Keyframe1                        int32
	Keyframe2                        int32
	Reconstruction_success_threshold float32
	Reconstruction_flag              int32
	Refine_camera_intrinsics         int16
	Pad2                             int16
	Dist                             float32
	Clean_frames                     int32
	Clean_action                     int32
	Clean_error                      float32
	Object_distance                  float32
	Pad3                             int32
}

// SDNA index: 494
type MovieTrackingStabilization struct {
	Flag      int32
	Tot_track int32
	Act_track int32
	Maxscale  float32
	Rot_track Pointer // *MovieTrackingTrack
	Locinf    float32
	Scaleinf  float32
	Rotinf    float32
	Filter    int32
	Ok        int32
	Scale     float32
}

// SDNA index: 495
type MovieTrackingReconstruction struct {
	Flag        int32
	Error       float32
	Last_camera int32
	Camnr       int32
	Cameras     Pointer // *MovieReconstructedCamera
}

// SDNA index: 496
type MovieTrackingObject struct {
	Next           Pointer // *MovieTrackingObject
	Prev           Pointer // *MovieTrackingObject
	Name           [64]int8
	Flag           int32
	Scale          float32
	Tracks         ListBase
	Reconstruction MovieTrackingReconstruction
	Keyframe1      int32
	Keyframe2      int32
}

// SDNA index: 497
type MovieTrackingStats struct {
	Message [256]int8
}

// SDNA index: 498
type MovieTrackingDopesheetChannel struct {
	Next         Pointer // *MovieTrackingDopesheetChannel
	Prev         Pointer // *MovieTrackingDopesheetChannel
	Track        Pointer // *MovieTrackingTrack
	Pad          int32
	Name         [64]int8
	Tot_segment  int32
	Segments     Pointer // *int32
	Max_segment  int32
	Total_frames int32
}

// SDNA index: 499
type MovieTrackingDopesheetCoverageSegment struct {
	Next        Pointer // *MovieTrackingDopesheetCoverageSegment
	Prev        Pointer // *MovieTrackingDopesheetCoverageSegment
	Coverage    int32
	Start_frame int32
	End_frame   int32
	Pad         int32
}

// SDNA index: 500
type MovieTrackingDopesheet struct {
	Ok                int32
	Sort_method       int16
	Flag              int16
	Coverage_segments ListBase
	Channels          ListBase
	Tot_channel       int32
	Pad               int32
}

// SDNA index: 501
type MovieTracking struct {
	Settings       MovieTrackingSettings
	Camera         MovieTrackingCamera
	Tracks         ListBase
	Reconstruction MovieTrackingReconstruction
	Stabilization  MovieTrackingStabilization
	Act_track      Pointer // *MovieTrackingTrack
	Objects        ListBase
	Objectnr       int32
	Tot_object     int32
	Stats          Pointer // *MovieTrackingStats
	Dopesheet      MovieTrackingDopesheet
}

// SDNA index: 502
type DynamicPaintSurface struct {
	Next                Pointer // *DynamicPaintSurface
	Prev                Pointer // *DynamicPaintSurface
	Canvas              Pointer // *DynamicPaintCanvasSettings
	Data                Pointer // *PaintSurfaceData
	Brush_group         Pointer // *Group
	Effector_weights    Pointer // *EffectorWeights
	Pointcache          Pointer // *PointCache
	Ptcaches            ListBase
	Current_frame       int32
	Name                [64]int8
	Format              int16
	Type                int16
	Disp_type           int16
	Image_fileformat    int16
	Effect_ui           int16
	Preview_id          int16
	Init_color_type     int16
	Pad_s               int16
	Flags               int32
	Effect              int32
	Image_resolution    int32
	Substeps            int32
	Start_frame         int32
	End_frame           int32
	Pad                 int32
	Init_color          [4]float32
	Init_texture        Pointer // *Tex
	Init_layername      [64]int8
	Dry_speed           int32
	Diss_speed          int32
	Color_dry_threshold float32
	Depth_clamp         float32
	Disp_factor         float32
	Spread_speed        float32
	Color_spread_speed  float32
	Shrink_speed        float32
	Drip_vel            float32
	Drip_acc            float32
	Influence_scale     float32
	Radius_scale        float32
	Wave_damping        float32
	Wave_speed          float32
	Wave_timescale      float32
	Wave_spring         float32
	Uvlayer_name        [64]int8
	Image_output_path   [1024]int8
	Output_name         [64]int8
	Output_name2        [64]int8
}

// SDNA index: 503
type DynamicPaintCanvasSettings struct {
	Pmd        Pointer // *DynamicPaintModifierData
	Dm         Pointer // *DerivedMesh
	Surfaces   ListBase
	Active_sur int16
	Flags      int16
	Pad        int32
	Error      [64]int8
}

// SDNA index: 504
type DynamicPaintBrushSettings struct {
	Pmd               Pointer // *DynamicPaintModifierData
	Dm                Pointer // *DerivedMesh
	Psys              Pointer // *ParticleSystem
	Mat               Pointer // *Material
	Flags             int32
	Collision         int32
	R                 float32
	G                 float32
	B                 float32
	Alpha             float32
	Wetness           float32
	Particle_radius   float32
	Particle_smooth   float32
	Paint_distance    float32
	Paint_ramp        Pointer // *ColorBand
	Vel_ramp          Pointer // *ColorBand
	Proximity_falloff int16
	Wave_type         int16
	Ray_dir           int16
	Pad               int16
	Wave_factor       float32
	Wave_clamp        float32
	Max_velocity      float32
	Smudge_strength   float32
}

// SDNA index: 505
type Mask struct {
	Id          ID
	Adt         Pointer // *AnimData
	Masklayers  ListBase
	Masklay_act int32
	Masklay_tot int32
	Sfra        int32
	Efra        int32
	Flag        int32
	Pad         int32
}

// SDNA index: 506
type MaskParent struct {
	Pad         int32
	Id_type     int32
	Id          Pointer // *ID
	Parent      [64]int8
	Sub_parent  [64]int8
	Parent_orig [2]float32
}

// SDNA index: 507
type MaskSplinePointUW struct {
	U    float32
	W    float32
	Flag int32
}

// SDNA index: 508
type MaskSplinePoint struct {
	Bezt   BezTriple
	Pad    int32
	Tot_uw int32
	Uw     Pointer // *MaskSplinePointUW
	Parent MaskParent
}

// SDNA index: 509
type MaskSpline struct {
	Next          Pointer // *MaskSpline
	Prev          Pointer // *MaskSpline
	Flag          int16
	Offset_mode   int8
	Weight_interp int8
	Tot_point     int32
	Points        Pointer // *MaskSplinePoint
	Parent        MaskParent
	Points_deform Pointer // *MaskSplinePoint
}

// SDNA index: 510
type MaskLayerShape struct {
	Next     Pointer // *MaskLayerShape
	Prev     Pointer // *MaskLayerShape
	Data     Pointer // *float32
	Tot_vert int32
	Frame    int32
	Flag     int8
	Pad      [7]int8
}

// SDNA index: 511
type MaskLayer struct {
	Next           Pointer // *MaskLayer
	Prev           Pointer // *MaskLayer
	Name           [64]int8
	Splines        ListBase
	Splines_shapes ListBase
	Act_spline     Pointer // *MaskSpline
	Act_point      Pointer // *MaskSplinePoint
	Alpha          float32
	Blend          int8
	Blend_flag     int8
	Falloff        int8
	Pad            [7]int8
	Flag           int8
	Restrictflag   int8
}

// SDNA index: 512
type RigidBodyWorld struct {
	Effector_weights      Pointer // *EffectorWeights
	Group                 Pointer // *Group
	Objects               Pointer // **Object
	Constraints           Pointer // *Group
	Pad                   int32
	Ltime                 float32
	Pointcache            Pointer // *PointCache
	Ptcaches              ListBase
	Numbodies             int32
	Steps_per_second      int16
	Num_solver_iterations int16
	Flag                  int32
	Time_scale            float32
	Physics_world         Pointer // *struct{}
}

// SDNA index: 513
type RigidBodyOb struct {
	Physics_object   Pointer // *struct{}
	Physics_shape    Pointer // *struct{}
	Type             int16
	Shape            int16
	Flag             int32
	Col_groups       int32
	Pad              int32
	Mass             float32
	Friction         float32
	Restitution      float32
	Margin           float32
	Lin_damping      float32
	Ang_damping      float32
	Lin_sleep_thresh float32
	Ang_sleep_thresh float32
	Orn              [4]float32
	Pos              [3]float32
	Pad1             float32
}

// SDNA index: 514
type RigidBodyCon struct {
	Ob1                       Pointer // *Object
	Ob2                       Pointer // *Object
	Type                      int16
	Num_solver_iterations     int16
	Flag                      int32
	Breaking_threshold        float32
	Pad                       float32
	Limit_lin_x_lower         float32
	Limit_lin_x_upper         float32
	Limit_lin_y_lower         float32
	Limit_lin_y_upper         float32
	Limit_lin_z_lower         float32
	Limit_lin_z_upper         float32
	Limit_ang_x_lower         float32
	Limit_ang_x_upper         float32
	Limit_ang_y_lower         float32
	Limit_ang_y_upper         float32
	Limit_ang_z_lower         float32
	Limit_ang_z_upper         float32
	Spring_stiffness_x        float32
	Spring_stiffness_y        float32
	Spring_stiffness_z        float32
	Spring_damping_x          float32
	Spring_damping_y          float32
	Spring_damping_z          float32
	Motor_lin_target_velocity float32
	Motor_ang_target_velocity float32
	Motor_lin_max_impulse     float32
	Motor_ang_max_impulse     float32
	Physics_constraint        Pointer // *struct{}
}

// SDNA index: 515
type FreestyleLineSet struct {
	Next               Pointer // *FreestyleLineSet
	Prev               Pointer // *FreestyleLineSet
	Name               [64]int8
	Flags              int32
	Selection          int32
	Qi                 int16
	Pad1               int16
	Qi_start           int32
	Qi_end             int32
	Edge_types         int32
	Exclude_edge_types int32
	Pad2               int32
	Group              Pointer // *Group
	Linestyle          Pointer // *FreestyleLineStyle
}

// SDNA index: 516
type FreestyleModuleConfig struct {
	Next         Pointer // *FreestyleModuleConfig
	Prev         Pointer // *FreestyleModuleConfig
	Script       Pointer // *Text
	Is_displayed int16
	Pad          [3]int16
}

// SDNA index: 517
type FreestyleConfig struct {
	Modules              ListBase
	Mode                 int32
	Raycasting_algorithm int32
	Flags                int32
	Sphere_radius        float32
	Dkr_epsilon          float32
	Crease_angle         float32
	Linesets             ListBase
}

// SDNA index: 518
type LineStyleModifier struct {
	Next      Pointer // *LineStyleModifier
	Prev      Pointer // *LineStyleModifier
	Name      [64]int8
	Type      int32
	Influence float32
	Flags     int32
	Blend     int32
}

// SDNA index: 519
type LineStyleColorModifier_AlongStroke struct {
	Modifier   LineStyleModifier
	Color_ramp Pointer // *ColorBand
}

// SDNA index: 520
type LineStyleAlphaModifier_AlongStroke struct {
	Modifier LineStyleModifier
	Curve    Pointer // *CurveMapping
	Flags    int32
	Pad      int32
}

// SDNA index: 521
type LineStyleThicknessModifier_AlongStroke struct {
	Modifier  LineStyleModifier
	Curve     Pointer // *CurveMapping
	Flags     int32
	Value_min float32
	Value_max float32
	Pad       int32
}

// SDNA index: 522
type LineStyleColorModifier_DistanceFromCamera struct {
	Modifier   LineStyleModifier
	Color_ramp Pointer // *ColorBand
	Range_min  float32
	Range_max  float32
}

// SDNA index: 523
type LineStyleAlphaModifier_DistanceFromCamera struct {
	Modifier  LineStyleModifier
	Curve     Pointer // *CurveMapping
	Flags     int32
	Range_min float32
	Range_max float32
	Pad       int32
}

// SDNA index: 524
type LineStyleThicknessModifier_DistanceFromCamera struct {
	Modifier  LineStyleModifier
	Curve     Pointer // *CurveMapping
	Flags     int32
	Range_min float32
	Range_max float32
	Value_min float32
	Value_max float32
	Pad       int32
}

// SDNA index: 525
type LineStyleColorModifier_DistanceFromObject struct {
	Modifier   LineStyleModifier
	Target     Pointer // *Object
	Color_ramp Pointer // *ColorBand
	Range_min  float32
	Range_max  float32
}

// SDNA index: 526
type LineStyleAlphaModifier_DistanceFromObject struct {
	Modifier  LineStyleModifier
	Target    Pointer // *Object
	Curve     Pointer // *CurveMapping
	Flags     int32
	Range_min float32
	Range_max float32
	Pad       int32
}

// SDNA index: 527
type LineStyleThicknessModifier_DistanceFromObject struct {
	Modifier  LineStyleModifier
	Target    Pointer // *Object
	Curve     Pointer // *CurveMapping
	Flags     int32
	Range_min float32
	Range_max float32
	Value_min float32
	Value_max float32
	Pad       int32
}

// SDNA index: 528
type LineStyleColorModifier_Material struct {
	Modifier   LineStyleModifier
	Color_ramp Pointer // *ColorBand
	Flags      int32
	Mat_attr   int32
}

// SDNA index: 529
type LineStyleAlphaModifier_Material struct {
	Modifier LineStyleModifier
	Curve    Pointer // *CurveMapping
	Flags    int32
	Mat_attr int32
}

// SDNA index: 530
type LineStyleThicknessModifier_Material struct {
	Modifier  LineStyleModifier
	Curve     Pointer // *CurveMapping
	Flags     int32
	Value_min float32
	Value_max float32
	Mat_attr  int32
}

// SDNA index: 531
type LineStyleGeometryModifier_Sampling struct {
	Modifier LineStyleModifier
	Sampling float32
	Pad      int32
}

// SDNA index: 532
type LineStyleGeometryModifier_BezierCurve struct {
	Modifier LineStyleModifier
	Error    float32
	Pad      int32
}

// SDNA index: 533
type LineStyleGeometryModifier_SinusDisplacement struct {
	Modifier   LineStyleModifier
	Wavelength float32
	Amplitude  float32
	Phase      float32
	Pad        int32
}

// SDNA index: 534
type LineStyleGeometryModifier_SpatialNoise struct {
	Modifier  LineStyleModifier
	Amplitude float32
	Scale     float32
	Octaves   int32
	Flags     int32
}

// SDNA index: 535
type LineStyleGeometryModifier_PerlinNoise1D struct {
	Modifier  LineStyleModifier
	Frequency float32
	Amplitude float32
	Angle     float32
	Octaves   int32
	Seed      int32
	Pad1      int32
}

// SDNA index: 536
type LineStyleGeometryModifier_PerlinNoise2D struct {
	Modifier  LineStyleModifier
	Frequency float32
	Amplitude float32
	Angle     float32
	Octaves   int32
	Seed      int32
	Pad1      int32
}

// SDNA index: 537
type LineStyleGeometryModifier_BackboneStretcher struct {
	Modifier        LineStyleModifier
	Backbone_length float32
	Pad             int32
}

// SDNA index: 538
type LineStyleGeometryModifier_TipRemover struct {
	Modifier   LineStyleModifier
	Tip_length float32
	Pad        int32
}

// SDNA index: 539
type LineStyleGeometryModifier_Polygonalization struct {
	Modifier LineStyleModifier
	Error    float32
	Pad      int32
}

// SDNA index: 540
type LineStyleGeometryModifier_GuidingLines struct {
	Modifier LineStyleModifier
	Offset   float32
	Pad      int32
}

// SDNA index: 541
type LineStyleGeometryModifier_Blueprint struct {
	Modifier        LineStyleModifier
	Flags           int32
	Rounds          int32
	Backbone_length float32
	Random_radius   int32
	Random_center   int32
	Random_backbone int32
}

// SDNA index: 542
type LineStyleGeometryModifier_2DOffset struct {
	Modifier LineStyleModifier
	Start    float32
	End      float32
	X        float32
	Y        float32
}

// SDNA index: 543
type LineStyleGeometryModifier_2DTransform struct {
	Modifier LineStyleModifier
	Pivot    int32
	Scale_x  float32
	Scale_y  float32
	Angle    float32
	Pivot_u  float32
	Pivot_x  float32
	Pivot_y  float32
	Pad      int32
}

// SDNA index: 544
type LineStyleThicknessModifier_Calligraphy struct {
	Modifier      LineStyleModifier
	Min_thickness float32
	Max_thickness float32
	Orientation   float32
	Pad           int32
}

// SDNA index: 545
type FreestyleLineStyle struct {
	Id                  ID
	Adt                 Pointer // *AnimData
	R                   float32
	G                   float32
	B                   float32
	Alpha               float32
	Thickness           float32
	Thickness_position  int32
	Thickness_ratio     float32
	Flag                int32
	Caps                int32
	Chaining            int32
	Rounds              int32
	Split_length        float32
	Min_angle           float32
	Max_angle           float32
	Min_length          float32
	Max_length          float32
	Split_dash1         int16
	Split_gap1          int16
	Split_dash2         int16
	Split_gap2          int16
	Split_dash3         int16
	Split_gap3          int16
	Pad                 int32
	Dash1               int16
	Gap1                int16
	Dash2               int16
	Gap2                int16
	Dash3               int16
	Gap3                int16
	Panel               int32
	Color_modifiers     ListBase
	Alpha_modifiers     ListBase
	Thickness_modifiers ListBase
	Geometry_modifiers  ListBase
}

type FileData struct{}
type GPUTexture struct{}
type Anim struct{}
type RenderResult struct{}
type ImBuf struct{}
type VFontData struct{}
type GHash struct{}
type Path struct{}
type SelBox struct{}
type EditFont struct{}
type BMEditMesh struct{}
type Cloth struct{}
type BVHTree struct{}
type DerivedMesh struct{}
type BVHTreeFromMesh struct{}
type Ocean struct{}
type OceanCache struct{}
type SculptSession struct{}
type RNG struct{}
type PTCacheEdit struct{}
type BodyPoint struct{}
type BodySpring struct{}
type SBScratch struct{}
type SceneStats struct{}
type DagForest struct{}
type RenderInfo struct{}
type RenderEngine struct{}
type ViewDepths struct{}
type SmoothView3DStore struct{}
type WmTimer struct{}
type SmoothView2DStore struct{}
type FileList struct{}
type FileLayout struct{}
type PanelType struct{}
type UiLayout struct{}
type UiListType struct{}
type SpaceType struct{}
type ARegionType struct{}
type Particle struct{}
type BNodeSocketType struct{}
type BNodeType struct{}
type UiBlock struct{}
type BNodeTreeType struct{}
type StructRNA struct{}
type BNodeInstanceHash struct{}
type BNodeTreeExec struct{}
type ParticleCacheKey struct{}
type KDTree struct{}
type ParticleDrawData struct{}
type LinkNode struct{}
type WmEvent struct{}
type WmSubWindow struct{}
type WmGesture struct{}
type PointerRNA struct{}
type WmOperatorType struct{}
type FLUID_3D struct{}
type WTURBULENCE struct{}
type MovieClipCache struct{}
type PaintSurfaceData struct{}
