use std::fs::File;
use std::io::copy;
use reqwest::blocking::get;

pub fn download(link: &str) -> Result<(), Box<dyn std::error::Error>> {
    let url = link;
    
    // Send GET request
    let response = get(url)?;
    
    let mut out = File::create("./images/download2.png")?;
    
    copy(&mut response.bytes()?.as_ref(), &mut out)?;
    
    println!("Downloaded file from {} to download.png", url);
    Ok(())

}