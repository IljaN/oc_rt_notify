<?php

namespace OCA\RealTimeNotifications\Authentication;


use Lcobucci\JWT\Builder;
use Lcobucci\JWT\Signer\Hmac\Sha512;

class TokenGenerator {

	private $secret;

	/** @var Sha512  */
	private $signer;

	/**
	 * @param string $secret
	 */
	public function __construct($secret) {

		if (!$secret || strlen(trim($secret)) < 15) {
			throw new \InvalidArgumentException(
				"Secret must have a length of min. 15 chars"
			);
		}

		$this->secret = $secret;
		$this->signer = new Sha512();
	}

	public function generateSubscriberToken($uid) {
		return (new Builder())
			->setNotBefore(time())
			->setIssuedAt(time())
			->setExpiration(time() + 3)
			->setAudience('subscriber')
			->setSubject($uid)
			->sign($this->signer, $this->secret)
			->getToken();
	}

	public function generatePublisherToken() {
		return (new Builder())
			->setNotBefore(time())
			->setIssuedAt(time())
			->setExpiration(time() + 3)
			->setAudience('publisher')
			->sign($this->signer, $this->secret)
			->getToken();
	}

}
