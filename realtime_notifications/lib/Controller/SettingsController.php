<?php

namespace OCA\RealTimeNotifications\Controller;

use OCA\RealTimeNotifications\EventRouterService;
use OCP\AppFramework\Controller;

class SettingsController extends Controller {

	/** @var EventRouterService  */
	private $er;

	public function __construct($appName, \OCP\IRequest $request, EventRouterService $er) {
		parent::__construct($appName, $request);
		$this->er = $er;
	}


	/**
	 * @NoAdminRequired
	 * @NoCSRFRequired
	 */
	public function get() {
		return new \OCP\AppFramework\Http\DataResponse(
			['backend_host' => $this->er->getBackendHost()]);
	}

}