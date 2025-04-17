use anchor_lang::prelude::*;

pub mod entrypoint {
    use super::*;
    declare_id!("F2dazhBjqX19HZBLvEHoToeggmDMj8xu3gaKkZ9YEpkU");
    pub use ID as ENTRYPOINT_ID;
}

pub mod onramp {
    use super::*;
    declare_id!("F4U8NchY73WXvMiLcfReLVUueXCtg9s5TGunmER3xLkJ");
    pub use ID as ONRAMP_ID;
}

pub mod offramp {
    use super::*;
    declare_id!("F3bceNrAhvhCmpbGx2HdBDzLpt91dQdn5yR8m7QwyiQC");
    pub use ID as OFFRAMP_ID;
}

pub mod compound {
    use super::*;
    declare_id!("F1G3cHLWhRtFQTnc7JaLNiEc5mFnkgDXo3Fobr8NcAiA");
    pub use ID as COMPOUND_ID;
}

pub mod token {
    use super::*;
    declare_id!("F5jMhWDPMrCbnCaw6mf91NtsRV4ccEuwGsrJUpBckQoD");
    pub use ID as TOKEN_ID;
}
