mod download_from_link;
mod stitch_content;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Example URL (replace with your image/video link)
    let url = "https://images.pexels.com/photos/6153343/pexels-photo-6153343.jpeg";

    // download_from_link::download(&url)?;
    stitch_content::stitching()?;
    Ok(())
}
