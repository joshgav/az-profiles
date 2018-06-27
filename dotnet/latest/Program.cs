using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using System.Linq;
using System.IO;
using Microsoft.Rest.Azure.Authentication;
using Microsoft.Azure.Management.ResourceManager;
using Microsoft.Azure.Management.ResourceManager.Models;
using dotenv.net;

namespace Profiles {
    public class Program {
		static string tenantID;
		static string clientID;
		static string clientSecret;
		static string subscriptionID;
		static string resourceGroupName = "dotnet-profiles-sample-01";
		static string location = "westus2";

        public static void Main(string[] args) {
            DotEnv.Config();

            tenantID = Environment.GetEnvironmentVariable("AZURE_TENANT_ID");
            clientID = Environment.GetEnvironmentVariable("AZURE_CLIENT_ID");
            clientSecret = Environment.GetEnvironmentVariable("AZURE_CLIENT_SECRET");
            subscriptionID = Environment.GetEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
			// assert all these are set

			CreateGroup(resourceGroupName).Wait();
        }

        public static async Task CreateGroup(string groupName)
        {
            // Build the service credentials and Azure Resource Manager clients
            var credentials = await ApplicationTokenProvider.LoginSilentAsync(tenantID, clientID, clientSecret);
            var resourceClient = new ResourceManagementClient(credentials);
            resourceClient.SubscriptionId = subscriptionID;

            Console.Write("creating group {0} in location {1}\n", groupName, location);
            var groupParams = new ResourceGroup { Location = location };
            resourceClient.ResourceGroups.CreateOrUpdate(groupName, groupParams);
            
            Console.Write("deleting group {0}\n", groupName);
            resourceClient.ResourceGroups.Delete(groupName);
        }
    }
}
