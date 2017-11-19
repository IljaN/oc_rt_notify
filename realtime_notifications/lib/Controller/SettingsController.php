<?php

namespace OCA\RealTimeNotifications\Controller;

use OCA\RealTimeNotifications\Config;
use OCP\AppFramework\Controller;
use OCP\AppFramework\Http\DataResponse;

class SettingsController extends Controller {

	/** @var Config  */
	private $cfg;

	public function __construct($appName, \OCP\IRequest $request, Config $cfg) {
		parent::__construct($appName, $request);
		$this->cfg = $cfg;
	}


	/**
	 * @NoAdminRequired
	 * @NoCSRFRequired
	 */
	public function getBackendHost() {
		return new DataResponse(
			['backend_host' => $this->cfg->getBackendHost()]);
	}

	/**
	 * @NoCSRFRequired
	 */
	public function getAdminSettings() {
		$resp = [
			'backend_host' => $this->cfg->getBackendHost(),
			'secret' => $this->cfg->getSecret()
		];

		return new DataResponse($resp);
	}

}