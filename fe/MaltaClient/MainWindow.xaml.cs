using Grpc.Core;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Forms;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using V1;
using static V1.MaltaBE;

namespace MaltaFE
{
    public partial class Main : Window
    {
        MaltaBEClient MaltaBEClient { get; set; }

        public Main()
        {
            InitializeComponent();

            Channel channel = new Channel("localhost:9090", ChannelCredentials.Insecure);
            MaltaBEClient = new MaltaBEClient(channel);
        }

        private void BrowseClick(object sender, RoutedEventArgs e)
        {
            OpenFileDialog fileDialog = new OpenFileDialog();
            
            if (fileDialog.ShowDialog() == System.Windows.Forms.DialogResult.OK)
            {
                FilePathText.Text = fileDialog.FileName;

                String hands = File.ReadAllText(fileDialog.FileName);

                CalculateRequest calculateRequest = new CalculateRequest();
                calculateRequest.Api = "1";
                calculateRequest.HackData = hands;

                if (AuthCheckbox.IsChecked.GetValueOrDefault())
                    calculateRequest.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJIYXBweVNhaWxhIn0.IIDVkUYHojF8Uy2xi3JtS8I5GBHX2vKNZNx8awegNt8";
                else
                    calculateRequest.Token = "";

                CalculateResponse calculateResponse = MaltaBEClient.Calculate(calculateRequest);
                OutputText.Text = calculateResponse.Data;
            }
        }

        private string ReplaceLastOccurrence(string Source, string Find, string Replace)
        {
            int place = Source.LastIndexOf(Find);

            if (place == -1)
                return Source;

            string result = Source.Remove(place, Find.Length).Insert(place, Replace);
            return result;
        }
    }
}
