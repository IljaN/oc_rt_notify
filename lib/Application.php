<?php
namespace OCA\RealTimeNotifications;


use OCA\RealTimeNotifications\Controller\SettingsController;
use OCP\AppFramework\App;
use OCP\AppFramework\IAppContainer;
use GuzzleHttp\Client as GuzzleClient;

class Application extends App {

	const APP_ID = 'realtime_notifications';

	/** @var EventRouterService  */
	private $eventRouter;



	/**
	 *
	 * @param array $urlParams
	 */
	public function __construct(array $urlParams = []) {

		parent::__construct(self::APP_ID, $urlParams);


		$c = $this->getContainer();
		$s = $c->getServer();

		$s->getEventDispatcher()->addListener(
			'OCA\Files::loadAdditionalScripts', function() {
				\OCP\Util::addScript(self::APP_ID, 'spop');
				\OCP\Util::addStyle(self::APP_ID, 'spop');
			}
		);



		$c->registerService('EventRouterService', function (IAppContainer $c) use ($s) {
			$er = new EventRouterService(
				$s->getConfig(),
				new GuzzleClient(),
				$s->getContentSecurityPolicyManager()
			);

			\OCP\Util::connectHook('OCP\Share', 'post_shared', $er, 'onPostShare');
			$er->run();

			return $er;
		});


		$this->eventRouter = $c->query('EventRouterService');

		$c->registerService('SettingsController', function (IAppContainer $c) use ($s) {
			return new SettingsController(
				self::APP_ID,
				$c->query('Request'),
				$this->eventRouter
			);
		});

	}

}