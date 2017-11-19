<?php

namespace OCA\RealTimeNotifications\Controller;


use OCA\RealTimeNotifications\Authentication\TokenGenerator;
use OCP\AppFramework\Controller;
use OCP\AppFramework\Http\DataResponse;
use OCP\IUserSession;


class AuthController extends Controller {

	/** @var IUserSession  */
	private $session;

	/** @var TokenGenerator  */
	private $tokenGenerator;

	/**
	 * AuthController constructor.
	 *
	 * @param string $appName
	 * @param \OCP\IRequest $request
	 * @param TokenGenerator $tokenGenerator
	 */
	public function __construct($appName, \OCP\IRequest $request, IUserSession $session, TokenGenerator $tokenGenerator) {
		parent::__construct($appName, $request);
		$this->session = $session;
		$this->tokenGenerator = $tokenGenerator;
	}

	/**
	 * @NoCSRFRequired
	 * @NoAdminRequired
	 */
	public function getAuthToken() {
		$token = $this->tokenGenerator->generateSubscriberToken(
			$this->session->getUser()->getUID()
		);
		return new DataResponse(['token' => sprintf($token)]);
	}

}