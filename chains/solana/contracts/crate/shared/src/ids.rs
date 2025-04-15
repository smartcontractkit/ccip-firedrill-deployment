use anchor_lang::prelude::*;

pub mod entrypoint {
    use super::*;
    declare_id!("Dri11c6TXJXELaZUgQtyw3doxbn7bGXQ6z5LRZ8bMPbc");
    pub use ID as ENTRYPOINT_ID;
}

pub mod onramp {
    use super::*;
    declare_id!("Dri11Hs4x2gXThgDwvjYoZ7xEPgfKNa5Ys2MWUqay3Tt");
    pub use ID as ONRAMP_ID;
}

pub mod offramp {
    use super::*;
    declare_id!("Dri11kiWR5f1oTn5NLm3wmfaDq3cYauR92wrsTxNaudt");
    pub use ID as OFFRAMP_ID;
}

pub mod compound {
    use super::*;
    declare_id!("Dri11NnBPbx85w8a47ZXBZa3Hi7E2xvHAA9Xpyk7UivQ");
    pub use ID as COMPOUND_ID;
}

pub mod token {
    use super::*;
    declare_id!("Dri11r4S1hY8iMpxNh6UoYZ9hsGgmGUUyZ3zxUb2rpc2");
    pub use ID as TOKEN_ID;
}
