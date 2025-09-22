use crate::module::shipping::usecase::ShippingUseCaseImpl;
use rdkafka::Message;
use rdkafka::message::BorrowedMessage;

pub struct Consumer {
    shipping_use_case: ShippingUseCaseImpl,
}

impl Consumer {
    pub fn new(shipping_use_case: ShippingUseCaseImpl) -> Self {
        Self { shipping_use_case }
    }

    pub async fn consume_snapshot_shippings_shipping_created(
        &self,
        message: BorrowedMessage<'_>,
    ) -> Result<(), anyhow::Error> {
        eprintln!("message {} ", message.topic());
        Ok(())
    }

    pub async fn consume_snapshot_shippings_shipping_updated(&self) -> Result<(), anyhow::Error> {
        Ok(())
    }
}
