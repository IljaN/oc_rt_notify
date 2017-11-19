<?php

namespace OCA\RealTimeNotifications\Authentication;


use Lcobucci\JWT\Builder;
use Lcobucci\JWT\Signer\Hmac\Sha512;

class TokenGenerator {


	private $secret;

	/** @var Sha512  */
	private $signer;

	/**
	 * AuthTokenGenerator constructor.
	 *
	 * @param string $secret
	 */
	public function __construct($secret) {
		$this->secret = $secret;
		$this->signer = new Sha512();
	}

	public function generateSubscriberToken($uid) {
		return (new Builder())
			->setNotBefore(time())
			->setIssuedAt(time())
			->setExpiration(time() + 10)
			->setAudience('subscriber')
			->setSubject($uid)
			->sign($this->signer, $this->secret)
			->getToken();
	}

}