use std::error::Error;
use std::process::Command;

pub fn stitching() -> Result<(), Box<dyn Error>> {
    // --- SETUP ---
    let image1_path = "./images/download.png";
    let image2_path = "./images/download2.png";
    let output_path = "./exports/output.mp4";
    let duration_per_image = "3";

    let output_width = "1080";
    let output_height = "1820";

    println!("Starting the image-to-video process...");

    // --- FILTER COMPLEX ---
    // scale + pad + setsar=1 to normalize SAR
    let filter_complex = format!(
        "[0:v]scale={w}:{h}:force_original_aspect_ratio=decrease,pad={w}:{h}:-1:-1:color=black,setsar=1[v0]; \
         [1:v]scale={w}:{h}:force_original_aspect_ratio=decrease,pad={w}:{h}:-1:-1:color=black,setsar=1[v1]; \
         [v0][v1]concat=n=2:v=1:a=0[outv]",
        w = output_width,
        h = output_height
    );

    // --- EXECUTE FFmpeg ---
    let status = Command::new("ffmpeg")
        .arg("-loop")
        .arg("1")
        .arg("-t")
        .arg(duration_per_image)
        .arg("-i")
        .arg(image1_path)
        .arg("-loop")
        .arg("1")
        .arg("-t")
        .arg(duration_per_image)
        .arg("-i")
        .arg(image2_path)
        .arg("-filter_complex")
        .arg(&filter_complex)
        .arg("-map")
        .arg("[outv]")
        .arg("-pix_fmt")
        .arg("yuv420p")
        .arg("-y")
        .arg(output_path)
        .status()?; // you can switch to .output()? to see stderr/stdout if needed

    if !status.success() {
        return Err("FFmpeg command failed to execute. Check console for details.".into());
    }

    println!("✅ Successfully created video '{}'", output_path);
    Ok(())
}
