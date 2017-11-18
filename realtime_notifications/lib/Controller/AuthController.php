<?php

namespace OCA\RealTimeNotifications\Controller;


use Lcobucci\JWT\Signer\Hmac\Sha512;
use OCP\AppFramework\Controller;
use OCP\AppFramework\Http\DataResponse;
use OCP\IUser;
use Lcobucci\JWT\Builder;
use OCP\IUserSession;


class AuthController extends Controller {

	/** @var IUserSession  */
	private $user;

	/** @var Sha512  */
	private $signer;

	/**
	 * AuthController constructor.
	 *
	 * @param string $appName
	 * @param \OCP\IRequest $request
	 * @param IUserSession $user
	 */
	public function __construct($appName, \OCP\IRequest $request, IUserSession $user) {
		parent::__construct($appName, $request);
		$this->signer = new Sha512();
		$this->user = $user;
	}

	/**
	 * @NoCSRFRequired
	 * @NoAdminRequired
	 */
	public function getAuthToken() {

		$token = (new Builder())
			->setNotBefore(time())
			->setIssuedAt(time())
			->setExpiration(time() + 10)
			->setAudience('subscriber')
			->setSubject($this->user->getUser()->getUID())
			->sign($this->signer, '186163c9826c3a0762319a81a3889dd9')
			->getToken();

		return new DataResponse(['token' => sprintf($token)]);
	}

}